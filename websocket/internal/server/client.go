package server

import (
	"github.com/gorilla/websocket"
	"time"
)

type Client struct {
	ClientId    int64
	Socket      *websocket.Conn
	ConnectTime uint64
	IsDeleted   bool
	StoreId     int64
	RoomId      int64
	Extend      string
	GroupList   []string
}

// NewClient
// @Auth：parker
// @Desc：初始化一个用户连接
// @Date：2024-05-08 14:11:30
// @param：clientId
// @param：conn
// @return：*Client
func NewClient(clientId int64, conn *websocket.Conn) *Client {
	return &Client{
		ClientId:    clientId,
		Socket:      conn,
		ConnectTime: uint64(time.Now().Unix()),
		IsDeleted:   false,
		Extend:      "",
		GroupList:   make([]string, 0),
	}
}
