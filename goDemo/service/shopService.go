package service

import (
	"encoding/json"
	"goDemo/entity"
	"goDemo/mysql"
	"goDemo/redis"
	"goDemo/utils"
	"log"
	"strconv"
	"time"

	Redis "github.com/go-redis/redis"
	"github.com/google/uuid"
)

func GetShopTest1(id string) string {
	result, err := redis.RedisClient.Get(utils.ShopIDPrefix + id).Result()
	if err == nil {
		return result
	}
	return ""
}

func GetShopTest2(id int) string {
	shop := &entity.Shop{}
	mysql.Db.Find(shop, "id = ?", id)
	jsons, _ := json.Marshal(shop)
	return string(jsons)
}

func GetShopByID(id int) (*entity.Shop, *utils.MyError) {
	shop := &entity.Shop{}
	myError := &utils.MyError{}
	result, err := redis.RedisClient.Get(utils.ShopIDPrefix + strconv.Itoa(id)).Result()
	if err == nil {
		json.Unmarshal([]byte(result), shop)
		return shop, nil
	}
	dbRes := mysql.Db.Find(shop, "id = ?", id)
	if dbRes.Error != nil {
		myError.Message1 = dbRes.Error.Error()
		myError.Message2 = utils.GetMysqlError
		return shop, myError
	}
	// 数据库找到，回写redis
	jsons, _ := json.Marshal(shop)
	err = redis.RedisClient.Set(utils.ShopIDPrefix+strconv.Itoa(id), jsons, utils.CommonExpireTime).Err()
	if err != nil {
		myError.Message1 = err.Error()
		myError.Message2 = utils.SetRedisError
		return shop, myError
	}
	return shop, nil
}

func GetShopListByTypeID(typeID int, field string) (*[]entity.Shop, *utils.MyError) {
	shopList := &[]entity.Shop{}
	myError := &utils.MyError{}
	result := mysql.Db.Where("type_id = ?", typeID).Order("sold").Select(field).Find(&shopList)
	if result.Error != nil {
		myError.Message1 = result.Error.Error()
		myError.Message2 = utils.GetMysqlError
		return shopList, myError
	}
	return shopList, nil
}

func GetShopListByPage(typeID, pageIndex, pageCount int, field string) (*[]entity.Shop, *utils.MyError) {
	shopList := &[]entity.Shop{}
	myError := &utils.MyError{}
	result := mysql.Db.Model(&entity.Shop{}).Where("type_id = ?", typeID).Limit(pageCount).Offset((pageIndex - 1) * pageCount).Select(field).Find(shopList)
	if result.Error != nil {
		myError.Message1 = result.Error.Error()
		myError.Message2 = utils.GetMysqlError
		return shopList, myError
	}
	return shopList, nil
}

func SaveShop(shop *entity.Shop) {
	mysql.Db.Model(shop).Create(shop)
}

func UpdateShopByID(id int, field, value string) *utils.MyError {
	myError := &utils.MyError{}
	// 查询数据库，未找到就返回
	result := mysql.Db.Model(&entity.Shop{}).Find(&entity.Shop{}, "id = ?", id)
	if result.Error != nil {
		myError.Message1 = result.Error.Error()
		myError.Message2 = utils.GetMysqlError
		return myError
	}
	// 先更新数据库
	result = mysql.Db.Model(&entity.Shop{}).Where("id = ?", id).Update(field, value)
	if result.Error != nil {
		myError.Message1 = result.Error.Error()
		myError.Message2 = utils.UpdateMysqlError
		return myError
	}
	// 再删除缓存
	redis.RedisClient.Del(utils.ShopIDPrefix + strconv.Itoa(id))
	return nil
}

// 缓存穿透：redis和mysql里都不存在数据，这样数据总会打到数据库
// 解决：缓存空数据
func GetShopByIDChuanTou(id int) (*entity.Shop, *utils.MyError) {
	shop := &entity.Shop{}
	myError := &utils.MyError{}

	result, err := redis.RedisClient.Get(utils.ShopIDPrefix + strconv.Itoa(id)).Result()
	if err == nil && len(result) > 0 {
		json.Unmarshal([]byte(result), shop)
		return shop, nil
	}
	// err != Redis.Nil说明能找到数据，但为空数据
	if len(result) == 0 && err != Redis.Nil {
		myError.Message2 = utils.ValueNullError
		return shop, myError
	}
	dbRes := mysql.Db.Find(shop, "id = ?", id)
	if dbRes.Error != nil {
		// 数据库没找到，缓存空数据
		redis.RedisClient.Set(utils.ShopIDPrefix+strconv.Itoa(id), "", utils.CommonExpireTime)

		myError.Message1 = dbRes.Error.Error()
		myError.Message2 = utils.ValueNullError
		return shop, myError
	}
	// 数据库找到，回写redis
	jsons, _ := json.Marshal(shop)
	err = redis.RedisClient.Set(utils.ShopIDPrefix+strconv.Itoa(id), jsons, utils.CommonExpireTime).Err()
	if err != nil {
		myError.Message1 = err.Error()
		myError.Message2 = utils.SetRedisError
		return shop, myError
	}
	return shop, nil
}

// 缓存雪崩：大量key同时失效或redis服务宕机，导致所有的请求都打到数据库
// 解决: 1. 大量key同时失效-->使用随机值作为过期时间，或者多级缓存
//       2. redis宕机-->使用redis集群

// 缓存击穿: 热点数据失效，所有的请求打到数据库
// 解决: 1. 热点数据永不过期 2. 若redis获取不到，则使用互斥锁来获取 3. 逻辑过期

// 下面是使用互斥锁来解决缓存击穿
func GetShopByIDJiChuan(id int) (*entity.Shop, *utils.MyError) {
	shop := &entity.Shop{}
	myError := &utils.MyError{}
	// 先从redis里获取
	key := utils.ShopIDPrefix + strconv.Itoa(id)
	result, err := redis.RedisClient.Get(key).Result()
	if err == nil && len(result) > 0 {
		// 在redis中能找到结果
		json.Unmarshal([]byte(result), shop)
		return shop, nil
	}
	if err == nil && len(result) == 0 {
		// 找到空字符串
		myError.Message2 = utils.ValueNullError
		return shop, myError
	}
	if err != Redis.Nil {
		// 并非没找到，而是查找的过程中出错了(例如服务宕机等)
		myError.Message1 = err.Error()
		myError.Message2 = utils.GetRedisError
		return shop, myError
	}
	// 在redis中没找到，尝试获取锁
	lockKey := utils.LockPrefix + strconv.Itoa(id)
	value := uuid.New().String()
	isLock := utils.TryLock(lockKey, value)
	if !isLock {
		time.Sleep(50 * time.Millisecond)
		return GetShopByIDJiChuan(id)
	}
	defer utils.UnLock(lockKey, value)
	// 获取到锁，查询数据库
	dbRes := mysql.Db.Model(&entity.Shop{}).Where("id = ?", id).Find(shop)
	if dbRes.RowsAffected == 0 {
		// 查询未错但没找到
		log.Println("未找到")
		redis.RedisClient.Set(key, "", utils.CommonExpireTime)
		myError.Message2 = utils.ValueNullError
		return shop, myError
	}
	if dbRes.Error != nil {
		// 查询出错
		myError.Message1 = dbRes.Error.Error()
		myError.Message2 = utils.GetMysqlError
		return shop, myError
	}
	// mysql找到，返回
	return shop, nil
}
