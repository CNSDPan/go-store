package server

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"store/common"
	"store/websocket/internal/types"
	"time"
)

type Subscribe struct {
	Log         logx.Logger
	PubSub      *redis.PubSub
	EnterPubSub *redis.PubSub
}

var SubWriteMsg = make(chan types.WriteMsg, 10000)
var SubWriteMsgBySingle = make(chan types.WriteMsg, 10000)

func NewSubscribe() (*Subscribe, error) {
	pubSub := AloneRedisClient.Subscribe(common.PubSubSocketMessageNormalChannelKey)
	//defer pubSub.Close()
	if _, err := pubSub.ReceiveTimeout(100 * time.Millisecond); err != nil {
		fmt.Printf("订阅 %s 接收消息异常，尝试 ping...", common.PubSubSocketMessageNormalChannelKey)
		if err = pubSub.Ping(""); err != nil {
			return &Subscribe{}, err
		}
	}

	enterPubSub := AloneRedisClient.Subscribe(common.PubSubSocketMessageLoginChannelKey)
	if _, err := enterPubSub.ReceiveTimeout(100 * time.Millisecond); err != nil {
		fmt.Printf("订阅 %s 接收消息异常，尝试 ping...", common.PubSubSocketMessageLoginChannelKey)
		if err = enterPubSub.Ping(""); err != nil {
			return &Subscribe{}, err
		}
	}
	return &Subscribe{PubSub: pubSub, EnterPubSub: enterPubSub}, nil
}

// SubReceive
// @Auth：
// @Desc：接收订阅的消息，并写入管道
// @Date：2024-06-15 17:48:42
// @receiver：sub
func (sub *Subscribe) SubReceive() {
	go func() {
		var (
			msg interface{}
			err error
		)
		defer sub.PubSub.Close()
		for {
			if msg, err = sub.PubSub.ReceiveTimeout(100 * time.Millisecond); err != nil {
				if err = sub.PubSub.Ping(""); err != nil {
					sub.Log.Info("订阅消息服务 PubSub.Ping timeout channel.name:%s", common.PubSubSocketMessageNormalChannelKey)
					break
				}
				continue
			}
			switch msg.(type) {
			case *redis.Message:
				m := msg.(*redis.Message)
				var writeMsg types.WriteMsg
				b := []byte(m.Payload)
				if err = jsonx.Unmarshal(b, &writeMsg); err != nil {
					sub.Log.Errorf("订阅消息服务 Receive Channel:%s json.Unmarshal  fail:%s", m.Channel, err.Error())
				} else {
					SubWriteMsg <- writeMsg
				}
			}
		}
		return
	}()
}

func (sub *Subscribe) EnterSubReceive() {
	go func() {
		var (
			msg interface{}
			err error
		)
		defer sub.EnterPubSub.Close()
		for {
			if msg, err = sub.EnterPubSub.ReceiveTimeout(100 * time.Millisecond); err != nil {
				if err = sub.EnterPubSub.Ping(""); err != nil {
					sub.Log.Info("订阅消息服务 PubSub.Ping timeout channel.name:%s", common.PubSubSocketMessageLoginChannelKey)
					break
				}
				continue
			}
			switch msg.(type) {
			case *redis.Message:
				m := msg.(*redis.Message)
				var writeMsg types.WriteMsg
				b := []byte(m.Payload)
				if err = jsonx.Unmarshal(b, &writeMsg); err != nil {
					sub.Log.Errorf("订阅消息服务 Receive Channel:%s json.Unmarshal  fail:%s", m.Channel, err.Error())
				} else {
					SubWriteMsgBySingle <- writeMsg
				}
			}
		}
		return
	}()
}
