package server

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"time"
)

const (
	// MaxMessageSize 消息大小
	MaxMessageSize = 8192
	// PingPeriod 每次ping的间隔时长
	PingPeriod = 30 * time.Second
	// PongPeriod 每次pong的间隔时长，可以是PingPeriod的一倍|两倍
	PongPeriod = 60 * time.Second
	// WriteWait client的写入等待时长
	WriteWait = 5 * time.Second
	// ReadWait client的读取等待时长
	ReadWait = 60 * time.Second
)

type Connect struct {
	ServerId string
	Node     *snowflake.Node
	logx.Logger
}

func NewConnect(serverId string, node *snowflake.Node) *Connect {
	return &Connect{
		ServerId: serverId,
		Node:     node,
	}
}

func (c *Connect) Run(w http.ResponseWriter, r *http.Request, webServer *Server) {
	c.Logger = logx.WithContext(r.Context())

	wsConn, err := (&websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		c.Logger.Errorf("websocket connect init fail:%v", err)
		http.NotFound(w, r)
		return
	}

	_ = wsConn.SetWriteDeadline(time.Now().Add(WriteWait))
	_ = wsConn.SetReadDeadline(time.Now().Add(ReadWait))
	wsConn.SetReadLimit(MaxMessageSize)
	wsConn.SetPongHandler(func(string) error {
		_ = wsConn.SetReadDeadline(time.Now().Add(PongPeriod))
		return nil
	})

	clientChannel := NewClient(wsConn)

	go webServer.readSubWriteMsg()
	go webServer.writeChannel(clientChannel)
	go webServer.readChannel(clientChannel)

}
