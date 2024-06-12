package server

import (
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"store/common"
	"store/websocket/internal/types"
)

type Subscribe struct {
	Log    logx.Logger
	PubSub *redis.PubSub
}

var SubWriteMsg = make(chan types.WriteMsg, 10000)

func NewSubscribe(log logx.Logger) (*Subscribe, error) {
	var (
		subscribe *Subscribe
		err       error
	)
	// 订阅redis的广播消息初始
	pubSub := AloneRedisClient.Subscribe(common.PubSubSocketMessageNormalChannelKey)
	_, err = pubSub.Receive()
	if err != nil {
		_ = pubSub.Close()
		return subscribe, err
	}
	subscribe.Log = log
	subscribe.PubSub = pubSub
	go subscribe.SubReceive()
	return subscribe, nil
}

// SubReceive
// @Auth：
// @Desc：接收订阅的消息，并写入管道
// @Date：2024-06-12 17:52:49
// @receiver：sub
func (sub *Subscribe) SubReceive() {
	var err error
	// 接收消息
	ch := sub.PubSub.ChannelSize(MaxMessageSize)
	for msg := range ch {
		var writeMsg types.WriteMsg
		b := []byte(msg.Payload)
		if err = jsonx.Unmarshal(b, &writeMsg); err != nil {
			sub.Log.Errorf("server subscribe Receive Channel:%s json.Unmarshal  fail:%s", msg.Channel, err.Error())
			continue
		}
		SubWriteMsg <- writeMsg
	}
}
