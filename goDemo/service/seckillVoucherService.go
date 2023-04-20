package service

import (
	"encoding/json"
	"fmt"
	"goDemo/entity"
	"goDemo/kafka"
	"goDemo/log2"
	"goDemo/mysql"
	"goDemo/redis"
	"goDemo/utils"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"

	Redis "github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var StockEnough = make(map[int]bool)
var SecKillExpireTime = make(map[int]time.Duration)

// 初始化，将库存加载到redis
// 同时计算出秒杀的时间间隔作为过期时间
func init() {
	vouchers := GetAllSeckillVoucher()
	if vouchers == nil {
		return
	}
	for _, voucher := range *vouchers {
		log.Println(voucher.Stock)
		redis.RedisClient.Set(utils.VoucherStockPrefix+strconv.Itoa(voucher.VoucherID), voucher.Stock, 0)
		StockEnough[voucher.VoucherID] = false
		begin := utils.StringToUnix(voucher.BeginTime)
		end := utils.StringToUnix(voucher.EndTime)
		SecKillExpireTime[voucher.VoucherID] = time.Duration(end - begin)
		log2.Info.Printf("过期时间: %d", end-begin)
	}
}

func GetAllSeckillVoucher() *[]entity.SecKillVoucher {
	vouchers := &[]entity.SecKillVoucher{}
	mysql.Db.Model(&entity.SecKillVoucher{}).Find(vouchers)
	return vouchers
}

// 最简单的秒杀
func SecKillSimple(voucherID, userID int) (*utils.MyError, int) {
	// 查询库存
	seckillVoucher := &entity.SecKillVoucher{}
	myError := &utils.MyError{}
	orderID := 0
	tx := mysql.Db.Begin()
	tx.Model(seckillVoucher).Where("voucher_id = ?", voucherID).Find(seckillVoucher)

	// 判断时间是否正常
	now := time.Now().Unix()
	if now < utils.StringToUnix(seckillVoucher.BeginTime) {
		myError.Message2 = utils.SecKillNotBegin
		return myError, orderID
	}
	if now > utils.StringToUnix(seckillVoucher.EndTime) {
		myError.Message2 = utils.SecKillAlreadyEnd
		return myError, orderID
	}
	// 库存不足
	if seckillVoucher.Stock <= 0 {
		myError.Message2 = utils.StockNotEnough
		return myError, orderID
	}
	// 更新库存
	res := tx.Exec("update tb_seckill_voucher set stock = stock-1 where voucher_id = ?", voucherID)
	//tx.Model(seckillVoucher).Where("voucher_id = ?", voucherID).Update("stock = stock-1")
	if res.Error != nil {
		myError.Message1 = res.Error.Error()
		myError.Message2 = utils.UpdateMysqlError
		return myError, orderID
	}
	if res.RowsAffected != 1 {
		myError.Message2 = utils.StockNotEnough
		return myError, orderID
	}
	// 更新库存成功才创建订单
	order := InitOrder(voucherID, userID)
	tx.Model(order).Create(order)
	tx.Commit()
	return nil, order.ID
}

// 乐观锁的秒杀
func SecKillWithOptimistic(voucherID, userID int) (*utils.MyError, int) {
	// 查询库存
	seckillVoucher := &entity.SecKillVoucher{}
	myError := &utils.MyError{}
	orderID := 0
	mysql.Db.Model(seckillVoucher).Where("voucher_id = ?", voucherID).Find(seckillVoucher)

	// 判断时间是否正常
	now := time.Now().Unix()
	if now < utils.StringToUnix(seckillVoucher.BeginTime) {
		myError.Message2 = utils.SecKillNotBegin
		return myError, orderID
	}
	if now > utils.StringToUnix(seckillVoucher.EndTime) {
		myError.Message2 = utils.SecKillAlreadyEnd
		return myError, orderID
	}
	// 库存不足
	if seckillVoucher.Stock <= 0 {
		myError.Message2 = utils.StockNotEnough
		return myError, orderID
	}
	// 使用乐观锁更新库存
	result := mysql.Db.Model(seckillVoucher).Where("voucher_id = ? and stock > 0", voucherID).
		Update("stock", gorm.Expr("stock - 1"))
	if result.Error != nil {
		myError.Message1 = result.Error.Error()
		myError.Message2 = utils.UpdateMysqlError
		return myError, orderID
	}
	if result.RowsAffected != 1 {
		myError.Message2 = utils.StockNotEnough
		return myError, orderID
	}
	// 更新库存成功才创建订单
	order := InitOrder(voucherID, userID)
	if err := mysql.Db.Create(order).Error; err != nil {
		myError.Message1 = err.Error()
		myError.Message2 = utils.OrderCreateError
		return myError, orderID
	}
	return nil, order.ID
}

// 一人一单的秒杀
func SecKillWithSingle(voucherID, userID int) (*utils.MyError, int) {
	// 查询库存
	seckillVoucher := &entity.SecKillVoucher{}
	myError := &utils.MyError{}
	orderID := 0
	mysql.Db.Model(seckillVoucher).Where("voucher_id = ?", voucherID).Find(seckillVoucher)

	// 判断时间是否正常
	now := time.Now().Unix()
	if now < utils.StringToUnix(seckillVoucher.BeginTime) {
		myError.Message2 = utils.SecKillNotBegin
		return myError, orderID
	}
	if now > utils.StringToUnix(seckillVoucher.EndTime) {
		myError.Message2 = utils.SecKillAlreadyEnd
		return myError, orderID
	}
	// 库存不足
	if seckillVoucher.Stock <= 0 {
		myError.Message2 = utils.StockNotEnough
		return myError, orderID
	}
	// 满足一人一单
	count := 0
	mysql.Db.Model(&entity.VoucherOrder{}).Where("voucher_id = ? and user_id = ?", voucherID, userID).Count(&count)
	if count >= 1 {
		myError.Message2 = utils.RepeatOrder
		return myError, orderID
	}

	// 使用乐观锁更新库存
	result := mysql.Db.Model(seckillVoucher).Where("voucher_id = ? and stock > 0", voucherID).
		Update("stock", gorm.Expr("stock - 1"))
	if result.Error != nil {
		myError.Message1 = result.Error.Error()
		myError.Message2 = utils.UpdateMysqlError
		return myError, orderID
	}
	if result.RowsAffected != 1 {
		myError.Message2 = utils.StockNotEnough
		return myError, orderID
	}
	// 更新库存成功才创建订单
	order := InitOrder(voucherID, userID)
	if err := mysql.Db.Create(order).Error; err != nil {
		myError.Message1 = err.Error()
		myError.Message2 = utils.OrderCreateError
		return myError, orderID
	}
	return nil, order.ID
}

// 一人一单加锁
func SecKillWithLockSingle(voucherID, userID int) (*utils.MyError, int) {
	// 查询库存
	seckillVoucher := &entity.SecKillVoucher{}
	myError := &utils.MyError{}
	orderID := 0
	mysql.Db.Model(seckillVoucher).Where("voucher_id = ?", voucherID).Find(seckillVoucher)

	// 判断时间是否正常
	now := time.Now().Unix()
	if now < utils.StringToUnix(seckillVoucher.BeginTime) {
		myError.Message2 = utils.SecKillNotBegin
		return myError, orderID
	}
	if now > utils.StringToUnix(seckillVoucher.EndTime) {
		myError.Message2 = utils.SecKillAlreadyEnd
		return myError, orderID
	}
	// 库存不足
	if seckillVoucher.Stock <= 0 {
		myError.Message2 = utils.StockNotEnough
		return myError, orderID
	}
	// 满足一人一单
	lockKey := utils.SecKillOrderLockPrefix + strconv.Itoa(userID)
	value := uuid.New().String()
	isLock := utils.TryLock(lockKey, value)
	if !isLock {
		myError.Message2 = utils.RepeatOrder
		return myError, orderID
	}
	if isLock {
		fmt.Println("获取到锁...")
		// time.Sleep(60 * time.Second)
		defer utils.UnLock(lockKey, value)
	}
	count := 0
	mysql.Db.Model(&entity.VoucherOrder{}).Where("voucher_id = ? and user_id = ?", voucherID, userID).Count(&count)
	if count >= 1 {
		myError.Message2 = utils.RepeatOrder
		return myError, orderID
	}

	// 使用乐观锁更新库存
	result := mysql.Db.Model(seckillVoucher).Where("voucher_id = ? and stock > 0", voucherID).
		Update("stock", gorm.Expr("stock - 1"))
	if result.Error != nil {
		myError.Message1 = result.Error.Error()
		myError.Message2 = utils.UpdateMysqlError
		return myError, orderID
	}
	if result.RowsAffected != 1 {
		myError.Message2 = utils.StockNotEnough
		return myError, orderID
	}
	// 更新库存成功才创建订单
	order := InitOrder(voucherID, userID)
	if err := mysql.Db.Create(order).Error; err != nil {
		myError.Message1 = err.Error()
		myError.Message2 = utils.OrderCreateError
		return myError, orderID
	}
	return nil, order.ID
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

func SecKillInit() {
	// 主要做的是将测试后新增的订单删除，方便重复测试
	// 1.将库存设置为50 2.将订单删除 3.将订单的自增设为1
	tx := mysql.Db.Begin()
	// id := 2
	// tx.Model(&entity.SecKillVoucher{}).Where("voucher_id = ?", id).Update("stock = ?", 50)
	tx.Exec("update tb_seckill_voucher set stock = 200 where voucher_id = 2")
	tx.Model(&entity.VoucherOrder{}).Where("id > 0").Delete(&entity.VoucherOrder{})
	tx.Exec("alter table tb_voucher_order auto_increment = 1")
	tx.Commit()
}

// TODO:引入redis和kafka

// 跳过一人一单(不好测试)，使用乐观锁和kafka
func SecKillWithKafkaV1(voucherID, userID int) *utils.MyError {

	// 查询库存
	seckillVoucher := &entity.SecKillVoucher{}
	myError := &utils.MyError{}

	mysql.Db.Model(seckillVoucher).Where("voucher_id = ?", voucherID).Find(seckillVoucher)

	// 判断时间是否正常
	now := time.Now().Unix()
	if now < utils.StringToUnix(seckillVoucher.BeginTime) {
		myError.Message2 = utils.SecKillNotBegin
		return myError
	}
	if now > utils.StringToUnix(seckillVoucher.EndTime) {
		myError.Message2 = utils.SecKillAlreadyEnd
		return myError
	}
	// 库存不足
	if seckillVoucher.Stock <= 0 {
		myError.Message2 = utils.StockNotEnough
		return myError
	}
	// 使用乐观锁更新库存
	result := mysql.Db.Model(seckillVoucher).Where("voucher_id = ? and stock > 0", voucherID).
		Update("stock", gorm.Expr("stock - 1"))
	if result.Error != nil {
		myError.Message1 = result.Error.Error()
		myError.Message2 = utils.UpdateMysqlError
		return myError
	}
	if result.RowsAffected != 1 {
		myError.Message2 = utils.StockNotEnough
		return myError
	}
	// 更新库存成功才创建订单
	order := InitOrder(voucherID, userID)
	orderJson, _ := json.Marshal(order)
	kafka.Send(string(orderJson))
	// if err := mysql.Db.Create(order).Error; err != nil {
	// 	myError.Message1 = err.Error()
	// 	myError.Message2 = utils.OrderCreateError
	// 	return myError, orderID
	// }
	return nil
}

// 引入redis，此处假定时间判断放在前端，只有时间符合才会存在秒杀按钮
func SecKillWithRedis(voucherID, userID int) *utils.MyError {
	myError := &utils.MyError{}
	// 查询redis标记
	value, ok := StockEnough[voucherID]
	if ok && value {
		myError.Message2 = utils.StockNotEnough
		return myError
	}
	userKey := utils.SecKillUserIDPrefix + strconv.Itoa(userID)
	// 一人一单
	/*
		这里的逻辑是数据库没数据，说明之前没有订单，就进行秒杀，
		秒杀成功了，将玩家id写入redis
		但这样有问题，高并发情况，当玩家还没被写入redis时
		玩家重复点击，两个请求都满足一人一单，就会出错
		解决办法1.是把这里到间隔的代码使用lua脚本
		2.是这里到间隔的代码加上用户锁
	*/
	_, err := redis.RedisClient.Get(userKey).Result()
	if err != Redis.Nil {
		myError.Message2 = utils.RepeatOrder
		return myError
	}

	// redis预减库存
	stockKey := utils.VoucherStockPrefix + strconv.Itoa(voucherID)
	result, err := redis.RedisClient.Decr(stockKey).Result()
	if err != nil {
		myError.Message2 = utils.GetRedisError
		return myError
	}
	if result < 0 {
		// 库存不足，进行标记
		StockEnough[voucherID] = true
		// 避免库存为负
		redis.RedisClient.Incr(stockKey)
		//
		myError.Message2 = utils.StockNotEnough
		return myError
	}
	if value, ok := SecKillExpireTime[voucherID]; ok {
		redis.RedisClient.Set(userKey, utils.CommonValue, value)
	}
	// ----间隔
	aux := &entity.SecKillAux{}
	aux.VoucherID = voucherID
	aux.UserID = userID
	jsons, _ := json.Marshal(aux)
	kafka.Send(string(jsons))
	return nil
}
