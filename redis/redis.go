package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // Redis 认证密码
		DB:       0,                // Redis 数据库索引
	})

	// 连接 Redis
	err := client.Ping().Err()
	if err != nil {
		panic(err)
	}

	// 设置键值对
	err = client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// 获取键的值
	val, err := client.Get("test").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("test:", val)

	// 关闭连接
	err = client.Close()
	if err != nil {
		panic(err)
	}
}
