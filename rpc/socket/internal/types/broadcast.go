package types

// WriteMsg
// @Auth：
// @Desc：广播消息结构
// @Date：2024-06-12 14:55:00
type WriteMsg struct {
	Version      int    `json:"version"`             // 用于区分业务版本号
	Operate      int    `json:"operate"`             // 操作
	Method       string `json:"method"`              // 事件
	SendRoomId   int64  `json:"sendRoomId"`          // 消息发送房间
	SendClientId int64  `json:"sendClientId,string"` // 消息发送指定人
	Extend       string `json:"extend"`              // 额外信息
	Body         []byte `json:"body"`                // 发送的主体内容
}
