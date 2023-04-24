package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

func main() {
	// 创建消费者配置
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	// 创建消费者组 A
	consumerGroupA, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "group-a", config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := consumerGroupA.Close(); err != nil {
			panic(err)
		}
	}()

	// 创建消费者组 B
	consumerGroupB, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "group-b", config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := consumerGroupB.Close(); err != nil {
			panic(err)
		}
	}()

	// 订阅主题
	topic := "my-topic"
	partitionList, err := consumerGroupA.Partitions(topic)
	if err != nil {
		panic(err)
	}

	// 创建信道，接收信号
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// 创建上下文，控制消费者运行状态
	ctx, cancel := context.WithCancel(context.Background())

	// 启动消费者协程
	for _, partition := range partitionList {
		go func(partition int32) {
			// 根据分区号选择消费者组
			consumerGroup := consumerGroupA
			if partition%2 == 0 {
				consumerGroup = consumerGroupB
			}

			partitionConsumer, err := consumerGroup.ConsumePartition(topic, partition, sarama.OffsetNewest)
			if err != nil {
				panic(err)
			}
			defer func() {
				if err := partitionConsumer.Close(); err != nil {
					panic(err)
				}
			}()
			for {
				select {
				case msg := <-partitionConsumer.Messages():
					fmt.Printf("Group: %s, Partition: %d, Offset: %d, Key: %s, Value: %s\n", consumerGroup.Config().ClientID, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				case <-signals:
					// 收到信号，停止消费者
					cancel()
					return
				case err := <-partitionConsumer.Errors():
					fmt.Printf("Group: %s, Partition: %d, Error: %v\n", consumerGroup.Config().ClientID, partition, err)
				}
			}
		}(partition)
	}

	// 等待上下文被取消
	<-ctx.Done()
	fmt.Println("Consumer stopped")
}

// func main() {
// 	ids := []int{5, 3, 4, 7}
// 	res := StringJoin(ids, ",")
// 	fmt.Println(res)
// }
