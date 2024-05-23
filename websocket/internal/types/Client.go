package types

import "github.com/gorilla/websocket"

type Client struct {
	ClientId    int64
	Websocket   *websocket.Conn
	ConnectTime uint64
	IsDeleted   bool
	StoreId     int64
	RoomId      int64
	Extend      string
	GroupList   []string
	Broadcast   chan Msg // 通过管道实时监控消息
}

// Msg
// @Auth：parker
// @Desc：服务端websocket推送到客户端的消息
// @Date：2024-05-23 11:08:16
type Msg struct {
	Version      int    `json:"version"`             // 用于区分业务版本号
	Operate      int    `json:"operate"`             // 操作
	Event        string `json:"event"`               // 事件
	SendClientId int64  `json:"sendClientId,string"` // 消息发送指定人
	Extend       string `json:"extend"`              // 额外信息
	Body         []byte `json:"body"`                // 发送的主体内容
}

// ReceiveMsg
// @Auth：parker
// @Desc：客户端websocket推送到服务端的消息
// @Date：2024-05-23 17:28:07
type ReceiveMsg struct {
	Version      int    `json:"version"`
	Operate      int    `json:"operate"`
	Event        string `json:"event"`
	StoreId      int64  `json:"storeId,string"`
	RoomId       int64  `json:"roomId,string"`
	FromClientId string `json:"fromClientId"`
	ToClientId   string `json:"toClientId"`
	Msg          string `json:"msg"`
}
