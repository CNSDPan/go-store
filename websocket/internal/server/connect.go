package server

import (
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

const (
	MaxMessageSize = 8192
)

type Connect struct {
	logx.Logger
}

func (c *Connect) Run(w http.ResponseWriter, r *http.Request) {
	ws, err := (&websocket.Upgrader{
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
	ws.SetReadLimit(MaxMessageSize)

	//c := &connection{}
}
