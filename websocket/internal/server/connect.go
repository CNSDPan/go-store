package server

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

const (
	MaxMessageSize = 8192
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
	wsConn.SetReadLimit(MaxMessageSize)

	clientId := c.Node.Generate().Int64()
	clientChannle := NewClient(clientId, wsConn)

	// 每个连接单独一个读写消息,避免消息拥挤
	go webServer.writeChannel(clientChannle)
	go webServer.readChannel(clientChannle)
}
