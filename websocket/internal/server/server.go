package server

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"store/tools"
	"store/websocket/internal/types"
	"strconv"
	"sync"
	"time"
)

const (
	// OperateSingleMsg 单人聊天操作
	OperateSingleMsg = 2
	// OperateGroupMsg 群体聊天操作
	OperateGroupMsg = 3
	// OperateConn 建立连接操作
	OperateConn = 10
)

type Server struct {
	Buckets       []*Bucket
	BucketIdx     uint32
	ServerId      string
	ClientManager ClientManager
	Log           logx.Logger
	Node          *snowflake.Node
}

// Bucket
// @Auth：parker
// @Desc：连接池
// @Date：2024-05-23 10:43:39
type Bucket struct {
	CLock   sync.RWMutex
	Clients map[int64]*Client
	Rooms   map[int64][]int64
	Idx     uint32
}

// NewServer
// @Auth：parker
// @Desc：初始化Server
// @Date：2024-05-14 14:32:44
// @return：*Server
func NewServer() *Server {
	return &Server{}
}

// StartWebsocket
// @Auth：parker
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
// @Auth：parker
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
// @Auth：parker
// @Desc：写消息的
// @Date：2024-05-22 15:47:30
// @receiver：s
func (s *Server) writeChannel(client *Client) {
	s.Log.Info("server websocket writeChannel start")
	s.Log.Infof("server websocket writeChannel PingPeriod:%v", PingPeriod)
	// ping前端的时隔
	ticker := time.NewTicker(PingPeriod)
	defer func() {
		ticker.Stop()
		_ = client.Websocket.Close()
	}()
	for {
		select {
		case message, ok := <-client.Broadcast:
			if !ok {
				s.Log.Error("server websocket client.Broadcast not ok ")
				_ = client.Websocket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := client.Websocket.NextWriter(websocket.TextMessage)
			if err != nil {
				s.Log.Errorf("server websocket.Conn.NextWriter fail:%s", err.Error())
				return
			}
			_, _ = w.Write(message.Body)
			if err = w.Close(); err != nil {
				s.Log.Errorf("server websocket w.Close() fail:%s", err.Error())
				return
			}
		case <-ticker.C:
			s.Log.Infof("server websocket writeChannel client.UserId:%d", client.UserId)
			// 心跳检测
			if err := client.Websocket.WriteMessage(websocket.PingMessage, nil); err != nil {
				s.Log.Errorf("server websocket.WriteMessage fail:%s", err.Error())
				return
			}
			//if err := client.Websocket.WriteMessage(websocket.TextMessage, []byte("自动发送")); err != nil {
			//	s.Log.Errorf("server websocket sand fail:%s", err.Error())
			//}
			s.Log.Infof("server websocket.WriteMessage ok")
		}
	}
}

// writeMessage
// @Auth：parker
// @Desc：读消息的
// @Date：2024-05-22 15:47:30
// @receiver：s
func (s *Server) readChannel(client *Client) {
	defer func() {
		//移出连接池
		if client.RoomId == 0 || client.UserId == 0 {
			s.Log.Infof("server websocket.readChannel client.RoomId || client.UserId eq 0")
			_ = client.Websocket.Close()
			return
		}
		s.Log.Infof("server websocket client.UserId:%d disconnect", client.UserId)
		s.getBucket(client.RoomId).DelBucket(client)
		// 断连后需要处理其他业务请求grpc
		s.ClientManager.DisConnect()
		_ = client.Websocket.Close()
	}()

	for {
		_, message, err := client.Websocket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				s.Log.Errorf("server websocket.ReadMessage fail:%s", err.Error())
				return
			}
		}
		if message == nil {
			s.Log.Infof("server websocket.ReadMessage message is nil:%v", message)
			return
		}
		// 将body信息转换成结构
		var websocketMsg types.ReceiveMsg
		if err = jsonx.Unmarshal(message, &websocketMsg); err != nil {
			s.Log.Errorf("server websocket.message json.Unmarshal websocketMsg fail:%s", err.Error())
			continue
		}
		client.AutoToken = websocketMsg.AuthToken
		client.RoomId = websocketMsg.RoomId

		client.Websocket.SetReadLimit(MaxMessageSize)
		client.Websocket.SetReadDeadline(time.Now().Add(ReadWait))
		client.Websocket.SetPongHandler(func(string) error {
			client.Websocket.SetReadDeadline(time.Now().Add(PongPeriod))
			return nil
		})
		switch websocketMsg.Operate {
		case OperateSingleMsg:
		case OperateGroupMsg:
			s.Log.Infof("打印发送群消息%v", websocketMsg.Event)
		case OperateConn:
			// client与server建立websocket成功后，client推送一次操作事件Operate:10，server将其进行连接池分组
			client.UserId, client.ClientId = s.ClientManager.Connect(client.AutoToken)
			if client.UserId == 0 {
				s.Log.Errorf("server websocket ClientManager.Connect user undefined by token:%s", client.AutoToken)
				return
			}
			// 获取要加入加入池子
			bucket := s.getBucket(websocketMsg.RoomId)
			s.Log.Infof("server websocket autoToken:%s bucket.Idx:%d", client.AutoToken, bucket.Idx)
			bucket.putBucket(client, websocketMsg.RoomId)
			// 请求grpc广播消息，通知群有用户进入群聊
		}
	}
}
