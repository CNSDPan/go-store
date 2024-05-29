package types

// Msg
// @Auth：parker
// @Desc：服务端websocket推送到客户端的消息
// @Date：2024-05-23 11:08:16
type Msg struct {
	Version      int    `json:"version"`             // 用于区分业务版本号
	Operate      int    `json:"operate"`             // 操作
	Method       string `json:"method"`              // 事件
	SendClientId int64  `json:"sendClientId,string"` // 消息发送指定人
	Extend       string `json:"extend"`              // 额外信息
	Body         []byte `json:"body"`                // 发送的主体内容
}

// ReceiveMsg
// @Auth：parker
// @Desc：客户端websocket推送到服务端的消息
// @Date：2024-05-23 17:28:07
type ReceiveMsg struct {
	Version      int    `json:"version"`         // 用于区分业务版本号
	Operate      int    `json:"operate"`         // 操作
	Method       string `json:"method"`          // 事件
	Event        Event  `json:"event,omitempty"` // 请求&响应参数
	StoreId      int64  `json:"storeId,string"`  //
	RoomId       int64  `json:"roomId,string"`   //
	FromClientId string `json:"fromClientId"`    // 消息发送人
	ToClientId   string `json:"toClientId"`      // 消息发送指定人
	Msg          string `json:"msg"`             //
	Extend       string `json:"extend"`          // 额外信息
	AuthToken    string `json:"authToken,omitempty"`
}

type Event struct {
	Params interface{} `json:"params,omitempty"` // 请求参数
	Data   interface{} `json:"data"`             // 响应参数
}
