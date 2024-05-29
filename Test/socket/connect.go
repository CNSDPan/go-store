package socket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/jsonx"
	"log"
	"strconv"
	"time"
)

type TestClient struct {
	Conn        *websocket.Conn
	sendMsgChan chan string
	recvMsgChan chan string
	timeout     int
	isAlive     bool
}

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

func New(url string) (tClient *TestClient, err error) {
	var d *websocket.Dialer
	d.HandshakeTimeout = 30 * time.Second
	conn, res, err := d.Dial(url, nil)
	//conn,res,err:=websocket.DefaultDialer.Dial("ws://127.0.0.1:6999/ws", nil)
	if err != nil {
		log.Println("拨号失败:", res)
		return nil, err
	}
	return &TestClient{
		Conn:        conn,
		sendMsgChan: make(chan string, 100),
		recvMsgChan: make(chan string, 100),
		timeout:     30,
		isAlive:     true,
	}, nil
}

// Auth
// @Auth：parker
// @Desc：连接事件
// @Date：2024-05-29 11:35:42
// @receiver：t
// @return：error
func (t *TestClient) Auth(authToken string) error {
	msg := ReceiveMsg{
		Version:   1,
		Operate:   10,
		Method:    "client",
		RoomId:    1,
		Msg:       "",
		AuthToken: authToken,
	}
	b, err := jsonx.Marshal(msg)
	if err != nil {
		fmt.Println("jsonx.Marshal fail:", err.Error())
		return err
	}
	err = t.Conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		fmt.Println("t.Conn.WriteMessage fail:", err.Error())
		return err
	}
	return nil
}

func (t *TestClient) Send() {
	go func() {
		r := ReceiveMsg{
			Version: 1,
			Operate: 3,
			Method:  "Msg",
			RoomId:  1,
			Msg:     "普通消息",
			Event:   Event{Params: ""},
		}
		for {
			msg := <-t.sendMsgChan
			r.Event.Params = msg
			b, err := jsonx.Marshal(r)
			if err != nil {
				log.Println("jsonx.Marshal fail:", err.Error())
				continue
			}
			err = t.Conn.WriteMessage(websocket.TextMessage, b)
			if err != nil {
				log.Println("send t.Conn.WriteMessage fail:", err.Error())
				continue
			}
		}
	}()
}

func (t *TestClient) Read() {
	go func() {
		var (
			mType int
			err   error
			b     []byte
		)
		for {
			mType, b, err = t.Conn.ReadMessage()
			if err != nil {
				t.recvMsgChan <- "读取失败：" + err.Error()
				break
			}
			t.recvMsgChan <- "读取" + strconv.Itoa(mType) + string(b)
		}
	}()
}

func (t *TestClient) TestAuth() error {
	msg := Data{
		Ip:       "",
		User:     "8",
		From:     "",
		Type:     "login",
		Content:  "",
		UserList: nil,
	}
	b, err := jsonx.Marshal(msg)
	if err != nil {
		fmt.Println("jsonx.Marshal fail:", err.Error())
		return err
	}
	err = t.Conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		fmt.Println("t.Conn.WriteMessage fail:", err.Error())
		return err
	}
	return nil
}
