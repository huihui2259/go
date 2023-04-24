package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// 基于sarama第三方库开发的kafka client

func main() {
	Consumer()
}
func Producer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "mytest"
	msg.Value = sarama.StringEncoder("1111")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"9.135.34.52:9092"}, config)
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}

func Consumer() {
	// 消费者示例
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer([]string{"9.135.34.52:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partitionList, err := consumer.Partitions("mytest") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("mytest", int32(partition), -1)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()

		go test(pc, partition)

		defer pc.AsyncClose()
	}
	for {

	}

}

func test(pc sarama.PartitionConsumer, a int) {
	fmt.Printf("test...: %d\n", a)
	for {
		select {
		case msg := <-pc.Messages():
			fmt.Printf("partition: %d, received message: %s\n", a, string(msg.Value))
		case err := <-pc.Errors():
			fmt.Printf("error: %s\n", err.Error())
		}
	}
}

func Consumer1() {
	// 消费者示例
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	consumer, err := sarama.NewConsumer([]string{"9.135.34.52:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partitionList, err := consumer.Partitions("mytest") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}

	done := make(chan struct{})
	defer close(done)

	for _, partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition("mytest", int32(partition), -1)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()

		go func(pc sarama.PartitionConsumer, partition int32) {
			fmt.Printf("start consuming partition %d\n", partition)
			defer fmt.Printf("stop consuming partition %d\n", partition)

			for {
				select {
				case msg := <-pc.Messages():
					fmt.Printf("partition %d: received message: %s\n", partition, string(msg.Value))
				case err := <-pc.Errors():
					fmt.Printf("partition %d: error: %s\n", partition, err.Error())
				case <-done:
					return
				}
			}
		}(pc, partition)
	}
}
