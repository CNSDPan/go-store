package server

import (
	"github.com/gorilla/websocket"
	"store/websocket/internal/types"
	"time"
)

// NewClient
// @Auth：parker
// @Desc：初始化一个用户连接
// @Date：2024-05-08 14:11:30
// @param：clientId
// @param：conn
// @return：*Client
func NewClient(clientId int64, conn *websocket.Conn) *types.Client {
	return &types.Client{
		ClientId:    clientId,
		Websocket:   conn,
		ConnectTime: uint64(time.Now().Unix()),
		IsDeleted:   false,
		Extend:      "",
		GroupList:   make([]string, 0),
	}
}
