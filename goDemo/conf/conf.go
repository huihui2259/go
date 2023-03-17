package conf

import (
	"log"
)

type DBConfig struct {
	User     string
	Password string
	Host     string

	DBName string
}

type RedisConfig struct {
}

var Config *DBConfig
var RedisCfg *RedisConfig

// 初始化配置
func init() {
	Config = &DBConfig{"root", "123456", "localhost:3306", "user"}
	log.Println(Config)
}
