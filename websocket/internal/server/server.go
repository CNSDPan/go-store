package server

import (
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/jsonx"
	"net/http"
	"store/websocket/internal/types"
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

type Server types.Server

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

// NewBuckets
// @Auth：parker
// @Desc：初始化websocket连接池
// @Date：2024-05-14 14:11:27
// @param：bucketNumber 池子数量
// @return：[]*Bucket
func NewBuckets(bucketNumber uint) []*types.Bucket {
	buckets := make([]*types.Bucket, bucketNumber)
	for i := uint(0); i < bucketNumber; i++ {
		buckets[i] = &types.Bucket{
			//clients: make(map[int64]*Client),
			Stores: make(map[int64]*types.Store),
			Rooms:  make(map[int64]*types.Room),
		}
	}
	return buckets
}

// writeMessage
// @Auth：parker
// @Desc：写消息的
// @Date：2024-05-22 15:47:30
// @receiver：s
func (s *Server) writeChannel(client *types.Client) {
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
			// 心跳检测
			if err := client.Websocket.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				s.Log.Errorf("server websocket.WriteMessage fail:%s", err.Error())
				return
			}
		}
	}
}

// writeMessage
// @Auth：parker
// @Desc：读消息的
// @Date：2024-05-22 15:47:30
// @receiver：s
func (s *Server) readChannel(client *types.Client) {
	defer func() {

	}()

	for true {
		_, message, err := client.Websocket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				s.Log.Errorf("server websocket.ReadMessage fail:%s", err.Error())
				continue
			}
		}
		if message == nil {
			s.Log.Info("server websocket.ReadMessage message is nil")
			continue
		}
		// 将body信息转换成结构
		var websocketMsg types.ReceiveMsg
		if err = jsonx.Unmarshal(message, &websocketMsg); err != nil {
			s.Log.Errorf("server websocket,message json.Unmarshal websocketMsg fail:%s", err.Error())
			continue
		}

		switch websocketMsg.Operate {
		case OperateSingleMsg:
		case OperateGroupMsg:
		case OperateConn:
			// client与server建立websocket成功后，client推送一次操作事件Operate:10，server将其进行连接池分组

		}
	}
}
