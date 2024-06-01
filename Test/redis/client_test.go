package redis

import (
	"fmt"
	"store/redis_db"
	"testing"
	"time"
)

func TestNewAloneRedis(t *testing.T) {
	aloneRedisClient, err := redis_db.NewAloneRedis()
	if err != nil {
		panic(err.Error())
	}
	err = aloneRedisClient.Set(time.Now().Format("2006-01-02"), time.Now().Format("2006-01-02 15:04:05"), -1).Err()
	fmt.Printf("set fail:%v", err)
}
