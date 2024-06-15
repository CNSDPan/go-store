package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/jsonx"
	"store/common"
	"store/redis_db"
	"testing"
	"time"
)

type WriteMsg struct {
	Version      int    `json:"version"`             // 用于区分业务版本号
	Operate      int    `json:"operate"`             // 操作
	Method       string `json:"method"`              // 事件
	SendRoomId   int64  `json:"sendRoomId"`          // 消息发送房间
	SendClientId int64  `json:"sendClientId,string"` // 消息发送指定人
	Extend       string `json:"extend"`              // 额外信息
	Body         []byte `json:"body"`                // 发送的主体内容
}

func TestAloneRedisPub(t *testing.T) {
	aloneRedisClient, err := redis_db.NewAloneRedis()
	if err != nil {
		panic(err.Error())
	}
	for {
		select {
		case <-time.After(10 * time.Second):
			tt := time.Now().Format("2006-01-02 15:04:05")
			writeMsg := WriteMsg{
				Version:      1,
				Operate:      3,
				Method:       "Normal",
				SendRoomId:   1,
				SendClientId: 0,
				Extend:       "",
				Body:         []byte(tt),
			}
			if body, err := jsonx.Marshal(writeMsg); err != nil {
				fmt.Printf("err :%s \r\n ", err.Error())
			} else {
				aloneRedisClient.Publish(common.PubSubSocketMessageNormalChannelKey, string(body))
				fmt.Printf("%s \r\n", tt)
			}
		}
	}
}

func TestAloneRedisSub(t *testing.T) {
	aloneRedisClient, err := redis_db.NewAloneRedis()
	if err != nil {
		panic(err.Error())
	}
	pubSub := aloneRedisClient.Subscribe(common.PubSubSocketMessageNormalChannelKey)

	go func() {
		defer func() {
			_ = pubSub.Close()
			panic("停止")
		}()
		for {
			msg, e := pubSub.Receive()
			if e != nil {
				fmt.Printf("订阅消息服务 Receive fail:%s", e.Error())
				break
			}
			switch msg.(type) {
			case *redis.Message:
				m, ok := msg.(*redis.Message)
				if !ok {
					fmt.Printf("订阅消息服务 Receive msg interface to *redis.Message not ok msg:%v \r\n ", msg)
				} else {
					var writeMsg WriteMsg
					b := []byte(m.Payload)
					if err := jsonx.Unmarshal(b, &writeMsg); err != nil {
						fmt.Printf("订阅消息服务 Receive Channel:%s json.Unmarshal  fail:%s", m.Channel, err.Error())
					} else {
						fmt.Printf("订阅 Payload:%s Channel:%s \r\n ", string(writeMsg.Body), m.Channel)
					}

				}
			default:
				fmt.Printf("订阅消息服务 Receive msg type no *redis.Message msg:%v \r\n ", msg)
			}
		}
	}()
	select {}
}

func TestAloneRedisSub2(t *testing.T) {
	aloneRedisClient, err := redis_db.NewAloneRedis()
	if err != nil {
		panic(err.Error())
	}
	pubSub := aloneRedisClient.Subscribe(common.PubSubSocketMessageNormalChannelKey)

	go func() {
		defer func() {
			_ = pubSub.Close()
			panic("停止")
		}()
		for {
			msg, e := pubSub.Receive()
			if e != nil {
				fmt.Printf("订阅消息服务 Receive fail:%s", e.Error())
				break
			}
			switch msg.(type) {
			case *redis.Message:
				m, ok := msg.(*redis.Message)
				if !ok {
					fmt.Printf("订阅消息服务 Receive msg interface to *redis.Message not ok msg:%v \r\n ", msg)
				} else {
					var writeMsg WriteMsg
					b := []byte(m.Payload)
					if err := jsonx.Unmarshal(b, &writeMsg); err != nil {
						fmt.Printf("订阅消息服务 Receive Channel:%s json.Unmarshal  fail:%s", m.Channel, err.Error())
					} else {
						fmt.Printf("订阅222 Payload:%s Channel:%s \r\n ", string(writeMsg.Body), m.Channel)
					}

				}
			default:
				fmt.Printf("订阅消息服务 Receive msg type no *redis.Message msg:%v \r\n ", msg)
			}
		}
	}()
	select {}
}

func TestAloneRedisSub3(t *testing.T) {
	aloneRedisClient, err := redis_db.NewAloneRedis()
	if err != nil {
		panic(err.Error())
	}

	pubSub := aloneRedisClient.Subscribe(common.PubSubSocketMessageNormalChannelKey)
	defer pubSub.Close()
	if _, err := pubSub.Receive(); err != nil {
		fmt.Printf(err.Error())
		return
	}
	go func() {
		ch := pubSub.Channel()
		for msg := range ch {
			var writeMsg WriteMsg
			b := []byte(msg.Payload)
			if err := jsonx.Unmarshal(b, &writeMsg); err != nil {
				fmt.Printf("订阅消息服务 Receive Channel:%s json.Unmarshal  fail:%s", msg.Channel, err.Error())
			} else {
				fmt.Printf("订阅3333 Payload:%s Channel:%s \r\n ", string(writeMsg.Body), msg.Channel)
			}
		}
	}()

	select {}
}
