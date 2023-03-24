package conf

import (
	"log"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	DBName   string
}

type RedisConfig struct {
	Addr     string
	Password string
	DBIndex  int
	PoolSize int
}

var Config *DBConfig
var RedisCfg *RedisConfig

// 初始化配置
func init() {
	Config = &DBConfig{"root", "123456", "9.135.34.52:3306", "demo"}
	log.Println(Config)
	RedisCfg = &RedisConfig{"127.0.0.1:6379", "", 0, 10}
	log.Println(RedisCfg)
}
