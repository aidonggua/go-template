package framework

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-template/configuration"
)

type Redis struct {
	*redis.Client
}

var redisClient = &Redis{
	Client: newRedisClient(),
}

func RedisInstance() *Redis {
	return redisClient
}

func newRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", configuration.Configs["redis.host"], configuration.Configs["redis.port"]),
	})
	return rdb
}
