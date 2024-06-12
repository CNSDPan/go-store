package server

import (
	"github.com/gorilla/websocket"
	"store/websocket/internal/types"
	"time"
)

type Client struct {
	AutoToken   string
	ClientId    int64
	Websocket   *websocket.Conn
	ConnectTime uint64
	StoreId     int64
	RoomId      int64
	UserId      int64
	Extend      string
	BucketId    string
	Broadcast   chan types.WriteMsg // 通过管道实时监控消息
}

// NewClient
// @Auth：
// @Desc：初始化一个用户连接
// @Date：2024-05-08 14:11:30
// @param：clientId
// @param：conn
// @return：*Client
func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		Websocket:   conn,
		ConnectTime: uint64(time.Now().Unix()),
		Extend:      "",
		Broadcast:   make(chan types.WriteMsg, 10000),
	}
}
