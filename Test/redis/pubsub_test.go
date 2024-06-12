package redis

import (
	"fmt"
	"store/common"
	"store/redis_db"
	"testing"
	"time"
)

func TestAloneRedisPub(t *testing.T) {
	aloneRedisClient, err := redis_db.NewAloneRedis()
	if err != nil {
		panic(err.Error())
	}
	for i := 1; i <= 2; i++ {
		num := i * 10000
		go func() {
			num2 := num
			for {
				select {
				case <-time.After(time.Microsecond):
					num2++
					if num2 >= num+5000 {
						return
					}
					aloneRedisClient.Publish(common.PubSubSocketMessageNormalChannelKey, fmt.Sprintf("发布消息：%d", num2))
					//fmt.Printf("发布消息 %d", num)
				}
			}
		}()
	}
	select {}
}

func TestAloneRedisSub(t *testing.T) {
	aloneRedisClient, err := redis_db.NewAloneRedis()
	if err != nil {
		panic(err.Error())
	}
	pubSub := aloneRedisClient.Subscribe(common.PubSubSocketMessageNormalChannelKey)
	if _, err = pubSub.Receive(); err != nil {
		fmt.Println("pubSub.Receive fail:")
		_ = pubSub.Close()
		panic(err.Error())
	}
	ch := pubSub.Channel()
	//for {
	//	select {
	//	case msg := <-ch:
	//		fmt.Printf("msg:%v chann:%v stirng:%v \r\n ", msg.Payload, msg.Channel, msg.String())
	//	case <-time.After(5 * time.Second):
	//		e := pubSub.Ping()
	//		fmt.Printf("%s ping %v \r\n ", time.Now().Format("2006-01-02 15:04:05"), e)
	//	}
	//}

	for msg := range ch {
		fmt.Printf("msg:%v chann:%v stirng:%v \r\n ", msg.Payload, msg.Channel, msg.String())
	}
}
