package utils

import (
	"goDemo/redis"
	"time"
)

func TryLock(key string) bool {
	flag, err := redis.RedisClient.SetNX(key, "1", 10*time.Second).Result()
	if err != nil {
		return false
	}
	return flag
}

func UnLock(key string) {
	redis.RedisClient.Del(key)
}
