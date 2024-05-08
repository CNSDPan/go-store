package server

import (
	"github.com/gorilla/websocket"
	"time"
)

type Client struct {
	ClientId    string
	SystemId    string
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
// @param：systemId
// @param：conn
// @return：*Client
func NewClient(clientId string, systemId string, conn *websocket.Conn) *Client {
	return &Client{
		ClientId:    clientId,
		SystemId:    systemId,
		Socket:      conn,
		ConnectTime: uint64(time.Now().Unix()),
		IsDeleted:   false,
		StoreId:     0,
		RoomId:      0,
		Extend:      "",
		GroupList:   make([]string, 0),
	}
}
