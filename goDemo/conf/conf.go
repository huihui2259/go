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

type KafkaConfig struct {
	Addr  string
	Topic string
}

var Config *DBConfig
var RedisCfg *RedisConfig
var KafkaCfg *KafkaConfig

// 初始化配置
func init() {
	Config = &DBConfig{"root", "123456", "9.135.34.52:3306", "demo"}
	log.Println(Config)
	RedisCfg = &RedisConfig{"127.0.0.1:6379", "", 0, 10}
	log.Println(RedisCfg)
	KafkaCfg = &KafkaConfig{"9.135.34.52:9092", "mytest"}
	log.Println(KafkaCfg)
}
