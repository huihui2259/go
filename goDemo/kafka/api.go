package kafka

import (
	"encoding/json"
	"fmt"
	"goDemo/conf"
	"goDemo/entity"
	"goDemo/log2"
	"goDemo/mysql"
	"goDemo/utils"

	"github.com/Shopify/sarama"
	"github.com/jinzhu/gorm"
)

// 这里特定是在testTopic下，消费消息的时候也是使用特定的方法
// 将数据发送到kafka
func Send(content string) {
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = conf.KafkaCfg.Topic
	msg.Value = sarama.StringEncoder(content)

	_, _, err := KafkaProducer.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
}

// 这里其实还可以扩展，根据不同的topic有不同的消费方法
func Receive() {
	partitionList, err := KafkaConsumer.Partitions(conf.KafkaCfg.Topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := KafkaConsumer.ConsumePartition(conf.KafkaCfg.Topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				myError := &utils.MyError{}
				aux := &entity.SecKillAux{}
				json.Unmarshal(msg.Value, aux)
				log2.Info.Println(string(msg.Value))

				// 乐观锁减库存
				result := mysql.Db.Model(&entity.SecKillVoucher{}).Where("voucher_id = ? and stock > 0", aux.VoucherID).
					Update("stock", gorm.Expr("stock - 1"))
				if result.Error != nil {
					myError.Message1 = result.Error.Error()
					myError.Message2 = utils.UpdateMysqlError
					log2.Error.Println(myError.String())
					return
				}
				if result.RowsAffected != 1 {
					myError.Message2 = utils.StockNotEnough
					log2.Error.Println(myError.String())
					return
				}

				// 创建订单，写入数据库
				order := InitOrder(aux.VoucherID, aux.UserID)
				if err := mysql.Db.Create(order).Error; err != nil {
					myError.Message1 = err.Error()
					myError.Message2 = utils.InsertMysqlError
					log2.Error.Println(myError.String())
					return
				}
				// redis.RedisClient.Set("test11", msg.Value, utils.CommonExpireTime)
				// fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
		}(pc)
	}
	// 防止主goroutine退出，否则上方的go func就无法执行而退出
	fmt.Scanln()
}

func InitOrder(voucherID, userID int) *entity.VoucherOrder {
	// seckillVoucher := &entity.SecKillVoucher{}
	// mysql.Db.Model(seckillVoucher).Where("voucher_id = ?", voucherID).Find(seckillVoucher)
	// if seckillVoucher.Stock <= 0 {
	// 	return nil
	// }
	order := &entity.VoucherOrder{}
	order.VoucherID = voucherID
	order.UserID = userID
	order.PayType = 0
	order.Status = 1
	order.PayTime = utils.InitTime
	order.UseTime = utils.TimeFormat
	order.RefundTime = utils.TimeFormat
	order.CreateTime = utils.GetNowTimeString()
	order.UpdateTime = utils.GetNowTimeString()
	return order
}
