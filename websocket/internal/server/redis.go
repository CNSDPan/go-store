package server

import (
	"github.com/go-redis/redis"
	"store/redis_db"
)

var AloneRedisClient *redis.Client

// InitAloneRedis
// @Auth：
// @Desc：初始化单机Redis
// @Date：2024-06-01 16:38:50
// @return：error
func InitAloneRedis() error {
	aloneRedisClient, err := redis_db.NewAloneRedis()
	if err != nil {
		return err
	}
	AloneRedisClient = aloneRedisClient
	return nil
}
