package redis

import (
	"fmt"
	"store/common"
	"store/redis_db"
	"testing"
	"time"
)

func TestNewAloneRedisSet(t *testing.T) {
	aloneRedisClient, err := redis_db.NewAloneRedis()
	if err != nil {
		panic(err.Error())
	}
	err = aloneRedisClient.Set(time.Now().Format("2006-01-02"), time.Now().Format("2006-01-02 15:04:05"), -1).Err()
	fmt.Printf("set fail:%v", err)
}

func TestNewAloneRedisLPush(t *testing.T) {
	aloneRedisClient, err := redis_db.NewAloneRedis()
	if err != nil {
		panic(err.Error())
	}
	err = aloneRedisClient.LPush(common.PubSubSocketMessageNormalChannelKey, time.Now().Format("2006-01-02")).Err()
	fmt.Printf("set fail:%v", err)
}
