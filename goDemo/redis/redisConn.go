package redis

import (
	"goDemo/conf"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.RedisCfg.Addr,
		Password: conf.RedisCfg.Password,
		DB:       conf.RedisCfg.DBIndex,
		PoolSize: conf.RedisCfg.PoolSize,
	})
}
