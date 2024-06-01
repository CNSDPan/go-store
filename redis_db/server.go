package redis_db

import (
	"github.com/go-redis/redis"
	"store/yaml"
	"sync"
)

var syncLock sync.Mutex

var AloneRedisClientMap = map[string]*redis.Client{}

func NewAloneRedis() (*redis.Client, error) {
	var (
		redisClient *redis.Client
		ok          bool
		err         error
	)
	syncLock.Lock()
	defer func() {
		_, err = redisClient.Ping().Result()
	}()
	defer syncLock.Unlock()
	if redisClient, ok = AloneRedisClientMap[yaml.RedisConf.AloneRedisConf.Addr]; ok {
		return redisClient, err
	}
	redisClient = redis.NewClient(&redis.Options{
		Addr:     yaml.RedisConf.AloneRedisConf.Addr,
		Password: yaml.RedisConf.AloneRedisConf.Password,
		DB:       yaml.RedisConf.AloneRedisConf.DB,
	})
	AloneRedisClientMap[yaml.RedisConf.AloneRedisConf.Addr] = redisClient
	return AloneRedisClientMap[yaml.RedisConf.AloneRedisConf.Addr], err
}
