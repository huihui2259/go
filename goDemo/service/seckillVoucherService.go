package service

import (
	"goDemo/entity"
	"goDemo/mysql"
	"goDemo/utils"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

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
	isLock := utils.TryLock(lockKey)
	if !isLock {
		myError.Message2 = utils.RepeatOrder
		return myError, orderID
	}
	if isLock {
		defer utils.UnLock(lockKey)
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
	utils.UnLock(lockKey)
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
	tx.Exec("update tb_seckill_voucher set stock = 50 where voucher_id = 2")
	tx.Model(&entity.VoucherOrder{}).Where("id > 0").Delete(&entity.VoucherOrder{})
	tx.Exec("alter table tb_voucher_order auto_increment = 1")
	tx.Commit()
}

// TODO:引入redis和kafka
