package server

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"store/common"
	"store/tools"
	"store/websocket/internal/types"
	"strconv"
	"sync"
	"time"
)

var module = "websocket服务下的server"

type Server struct {
	Buckets       []*Bucket
	BucketIdx     uint32
	ServerId      string
	ClientManager ClientManager
	Log           logx.Logger
	Node          *snowflake.Node
}

// Bucket
// @Auth：
// @Desc：连接池
// @Date：2024-05-23 10:43:39
type Bucket struct {
	CLock   sync.RWMutex
	Clients map[int64]*Client
	Rooms   map[int64][]int64
	Idx     uint32
}

// NewServer
// @Auth：
// @Desc：初始化Server
// @Date：2024-05-14 14:32:44
// @return：*Server
func NewServer() *Server {
	return &Server{}
}

// StartWebsocket
// @Auth：
// @Desc：启动websocket服务
// @Date：2024-05-14 14:34:52
// @receiver：s
func (s *Server) StartWebsocket() {
	connect := NewConnect(s.ServerId, s.Node)
	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		connect.Run(writer, request, s)
	})
}

// getBucket
// @Auth：
// @Desc：通过群聊房间ID得出所在连接池
// @Date：2024-05-28 14:24:13
// @receiver：s
// @param：roomId
func (s *Server) getBucket(roomId int64) *Bucket {
	roomIdStr := strconv.FormatInt(roomId, 10)
	// 通过cityHash算法 % 池子数量进行取模,得出需要放入哪个连接池里
	idx := tools.CityHash32([]byte(roomIdStr), uint32(len(roomIdStr))) % s.BucketIdx
	return s.Buckets[idx]
}

// writeMessage
// @Auth：
// @Desc：写消息的
// @Date：2024-05-22 15:47:30
// @receiver：s
func (s *Server) writeChannel(client *Client) {
	// ping前端的时隔
	ticker := time.NewTicker(PingPeriod)
	defer func() {
		ticker.Stop()
		_ = client.Websocket.Close()
	}()
	for {
		select {
		case message, ok := <-client.Broadcast:
			// 每次写之前，都需要设置超时时间，如果只设置一次就会出现总是超时
			_ = client.Websocket.SetWriteDeadline(time.Now().Add(WriteWait))
			if !ok {
				s.Log.Errorf("%s writeChannel <- client.Broadcast not ok ", module)
				_ = client.Websocket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := client.Websocket.NextWriter(websocket.TextMessage)
			if err != nil {
				s.Log.Errorf("%s Conn.NextWriter fail:%s", module, err.Error())
				return
			}
			_, _ = w.Write(message.Body)
			if err = w.Close(); err != nil {
				s.Log.Errorf("%s w.Close() fail:%s", module, err.Error())
				return
			}
		case <-ticker.C:
			// 每次写之前，都需要设置超时时间，如果只设置一次就会出现总是超时
			client.Websocket.SetWriteDeadline(time.Now().Add(PingPeriod))
			// 心跳检测
			if err := client.Websocket.WriteMessage(websocket.PingMessage, nil); err != nil {
				s.Log.Errorf("%s WriteMessage fail:%s", module, err.Error())
				return
			}
		}
	}
}

// writeMessage
// @Auth：
// @Desc：读消息的
// @Date：2024-05-22 15:47:30
// @receiver：s
func (s *Server) readChannel(client *Client) {
	var (
		methodCode string
		methodMsg  string
		methodErr  error
	)
	defer func() {
		//移出连接池
		if client.RoomId == 0 || client.UserId == 0 {
			s.Log.Infof("%s readChannel client.RoomId || client.UserId eq 0", module)
			_ = client.Websocket.Close()
			return
		}
		s.Log.Infof("%s client.UserId:%d RoomId:%d disconnect", module, client.UserId, client.RoomId)
		s.getBucket(client.RoomId).DelBucket(client)
		s.Log.Infof("%s disconnect 数量:%d", module, len(s.getBucket(client.RoomId).Rooms[client.RoomId]))
		// 断连后需要处理其他业务请求grpc
		s.ClientManager.DisConnect()
		_ = client.Websocket.Close()
	}()

	for {
		messageType, message, err := client.Websocket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				s.Log.Errorf("%s ReadMessage fail:%s", module, err.Error())
				return
			}
		}
		// 断连的messageType为-1
		if message == nil && messageType == -1 {
			s.Log.Infof("%s ReadMessage message UserId:%d is nil:%v messageType:%d", module, client.UserId, message, messageType)
			return
		}
		// 将body信息转换成结构
		var websocketMsg types.ReceiveMsg
		if err = jsonx.Unmarshal(message, &websocketMsg); err != nil {
			s.Log.Errorf("%s message json.Unmarshal websocketMsg fail:%s", module, err.Error())
			continue
		}

		if websocketMsg.FromClientId == "" {
			websocketMsg.FromClientId = "0"
		}
		if websocketMsg.ToClientId == "" {
			websocketMsg.ToClientId = "0"
		}
		client.AutoToken = websocketMsg.AuthToken
		client.RoomId = websocketMsg.RoomId
		if client.ClientId != 0 {
			websocketMsg.FromClientId = strconv.FormatInt(client.ClientId, 10)
			websocketMsg.FromUserName = client.Name
		}
		client.Websocket.SetReadLimit(MaxMessageSize)
		_ = client.Websocket.SetReadDeadline(time.Now().Add(ReadWait))
		client.Websocket.SetPongHandler(func(string) error {
			_ = client.Websocket.SetReadDeadline(time.Now().Add(PongPeriod))
			return nil
		})

		switch websocketMsg.Operate {
		case tools.OPERATE_SINGLE_MSG:
		case tools.OPERATE_GROUP_MSG:
			methodCode, methodMsg, methodErr = s.ClientManager.MethodHandle(websocketMsg, s.Log)

		case tools.OPERATE_CONN_MSG:
			// client与server建立websocket成功后，client推送一次操作事件Operate:10，server将其进行连接池分组
			client.UserId, client.ClientId, client.Name = s.ClientManager.Connect(client.AutoToken)
			if client.UserId == 0 {
				s.Log.Errorf("%s ClientManager.Connect user undefined by token:%s", module, client.AutoToken)
				return
			}
			// 获取要加入加入池子
			bucket := s.getBucket(websocketMsg.RoomId)
			bucket.putBucket(client, websocketMsg.RoomId)
			// 请求grpc广播消息，通知群有用户进入群聊
			s.Log.Infof(
				"%s 进群了:Idx:%v、rooms.len:%v",
				client.Name, bucket.Idx, len(bucket.Rooms[client.RoomId]),
			)
			websocketMsg.FromClientId = strconv.FormatInt(client.ClientId, 10)
			websocketMsg.FromUserName = client.Name
			methodCode, methodMsg, methodErr = s.ClientManager.MethodHandle(websocketMsg, s.Log)
		}
		if methodCode != common.RESPONSE_SUCCESS {
			s.Log.Errorf("%s 广播消息 methodCode:%s methodMsg:%s", methodCode, methodMsg)
		}
		if methodErr != nil {
			s.Log.Errorf("读消息管道：echo websocketMsg.Operate:%d methodErr:%s", websocketMsg.Operate, methodErr.Error())
		}
	}
}

// readSubWriteMsg
// @Auth：
// @Desc：处理消息，并发送房间内的人
// @Date：2024-06-12 18:05:19
// @receiver：s
func (s *Server) readSubWriteMsg() {
	for {
		select {
		case writeMsg := <-SubWriteMsg:
			b := s.getBucket(writeMsg.SendRoomId)
			clients := b.Rooms[writeMsg.SendRoomId]
			for _, clientId := range clients {
				client, ok := b.Clients[clientId]
				if ok {
					client.Broadcast <- writeMsg
				}
			}
		}
	}
}

// readSubWriteMsgBySingle
// @Auth：
// @Desc：处理消息，并发送给房间的指定人
// @Date：2024-06-25 11:28:00
// @receiver：s
func (s *Server) readSubWriteMsgBySingle() {
	for {
		select {
		case writeMsg := <-SubWriteMsgBySingle:
			b := s.getBucket(writeMsg.SendRoomId)
			client, ok := b.Clients[writeMsg.SendClientId]
			if ok {
				client.Broadcast <- writeMsg
			}
		}
	}
}
