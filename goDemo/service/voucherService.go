package service

import (
	"encoding/json"
	"goDemo/entity"
	"goDemo/mysql"
	"goDemo/redis"
	"goDemo/utils"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetVoucherListByShopID(id int) (*[]entity.Voucher, *utils.MyError) {
	voucherList := &[]entity.Voucher{}
	myError := &utils.MyError{}
	err := mysql.Db.Model(&entity.Voucher{}).Where("shop_id = ?", id).Find(voucherList).Error
	if err != nil {
		myError.Message1 = err.Error()
		myError.Message2 = utils.GetMysqlError
		return voucherList, myError
	}
	return voucherList, nil
}

func GetVoucherByID(id int) (*entity.Voucher, *utils.MyError) {
	voucher := &entity.Voucher{}
	myError := &utils.MyError{}
	// 首先查询redis
	key := utils.VoucherPrefix + strconv.Itoa(id)
	result, err := redis.RedisClient.Get(key).Result()
	if err == nil && result != "" {
		// redis 找到
		json.Unmarshal([]byte(result), voucher)
		return voucher, nil
	}
	if err == nil && result == "" {
		// redis找到空数据
		myError.Message1 = err.Error()
		myError.Message2 = utils.ValueNullError
		return voucher, myError
	}
	// redis未找到
	err = mysql.Db.Model(&entity.Voucher{}).Where("id = ?", id).Find(voucher).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 数据库没找到数据，缓存空数据
			redis.RedisClient.Set(key, "", utils.CommonExpireTime)
		}
		myError.Message1 = err.Error()
		myError.Message2 = utils.GetMysqlError
		return voucher, myError
	}
	// mysql找到数据，缓存到redis
	jsons, _ := json.Marshal(voucher)
	redis.RedisClient.Set(key, jsons, utils.CommonExpireTime)
	return voucher, myError
}

func AddSeckillVoucher(voucher *entity.Voucher) *utils.MyError {
	myError := &utils.MyError{}
	tx := mysql.Db.Begin()
	err := tx.Create(voucher).Error
	if err != nil {
		myError.Message1 = err.Error()
		myError.Message2 = utils.InsertMysqlError
		return myError
	}
	key := utils.VoucherPrefix + strconv.Itoa(voucher.ID)
	jsons, _ := json.Marshal(voucher)
	redis.RedisClient.Set(key, string(jsons), utils.CommonExpireTime)
	if voucher.Type == 0 {
		return nil
	}
	seckillVoucher := &entity.SecKillVoucher{}
	seckillVoucher.VoucherID = voucher.ID
	seckillVoucher.Stock = voucher.Stock
	seckillVoucher.CreateTime = voucher.CreateTime
	seckillVoucher.UpdateTime = voucher.UpdateTime
	seckillVoucher.BeginTime = voucher.BeginTime
	seckillVoucher.EndTime = voucher.EndTime
	a, _ := json.Marshal(seckillVoucher)
	log.Println(string(a))
	err = tx.Create(seckillVoucher).Error
	if err != nil {
		myError.Message1 = err.Error()
		myError.Message2 = utils.InsertMysqlError
		return myError
	}
	secKey := utils.SecKillVoucherPrefix + strconv.Itoa(voucher.ID)
	secKill, _ := json.Marshal(seckillVoucher)
	redis.RedisClient.Set(secKey, string(secKill), utils.CommonExpireTime)
	tx.Commit()
	return nil
}
