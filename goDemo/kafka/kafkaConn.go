package kafka

import (
	"fmt"
	"goDemo/conf"

	"github.com/Shopify/sarama"
)

var KafkaProducer sarama.SyncProducer
var KafkaConsumer sarama.Consumer
var Topic = "mytest"

// 初始化kafka
func init() {
	var err error
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	KafkaProducer, err = sarama.NewSyncProducer([]string{conf.KafkaCfg.Addr}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	KafkaConsumer, err = sarama.NewConsumer([]string{conf.KafkaCfg.Addr}, config)
	if err != nil {
		fmt.Println("consumer closed, err:", err)
		return
	}
}
