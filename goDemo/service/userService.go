package service

import (
	"encoding/json"
	"goDemo/entity"
	"goDemo/mysql"
	"goDemo/redis"
	"goDemo/utils"
	"log"
	"strconv"
)

func GetUserByID(id int) (*entity.User, *utils.MyError) {
	user := &entity.User{}
	myError := &utils.MyError{}
	// 先查询redis
	result, err := redis.RedisClient.Get(strconv.Itoa(id)).Result()
	if err == nil {
		// redis能找到
		json.Unmarshal([]byte(result), user)
		return user, nil
	}
	// redis未找到，查询mysql
	dbResult := mysql.Db.Find(user, "id = ?", id)
	log.Println(user)
	if dbResult.Error != nil {
		// mysql未找到，返回错误
		myError.Message1 = dbResult.Error.Error()
		myError.Message2 = utils.GetMysqlError
		return user, myError
	}
	// mysql找到，就再次插入redis
	jsons, _ := json.Marshal(user)
	if err := redis.RedisClient.Set(utils.UserIDPrefix+strconv.Itoa(id), string(jsons), 0).Err(); err != nil {
		myError.Message1 = err.Error()
		myError.Message2 = utils.SetRedisError
		return user, myError
	}
	return user, nil
}

func RegisterUser(user *entity.User) *utils.MyError {
	log.Println(user)
	myError := &utils.MyError{}
	// 这里以电话号码作为唯一标识
	result := mysql.Db.Create(user)
	if result.Error != nil {
		myError.Message1 = result.Error.Error()
		myError.Message2 = utils.InsertMysqlError
		return myError
	}
	// 注册成功，写入redis
	jsons, _ := json.Marshal(user)
	if err := redis.RedisClient.Set(utils.UserIDPrefix+strconv.Itoa(user.ID), string(jsons), 0).Err(); err != nil {
		myError.Message1 = err.Error()
		myError.Message2 = utils.SetRedisError
		return myError
	}
	return nil
}

func Delete(id int) *utils.MyError {
	myError := &utils.MyError{}
	result := mysql.Db.Where("id = ?", id).Delete(&entity.User{})
	if result.Error != nil {
		myError.Message1 = result.Error.Error()
		myError.Message2 = utils.DeleteMysqlError
		return myError
	}
	return nil
}

func UpdateNickName(user *entity.User) *utils.MyError {
	myError := &utils.MyError{}
	var selectUser *entity.User
	var err *utils.MyError
	// 若不存在，那就更改错误
	if selectUser, err = GetUserByID(user.ID); err != nil {
		return err
	}
	if selectUser.NickName == user.NickName {
		myError.Message2 = "修改后的昵称与修改前相同"
		return myError
	}
	result := mysql.Db.Model(user).Where("id = ?", user.ID).Update("update_time", user.UpdateTime).Update("nick_name", user.NickName)
	if result.Error != nil {
		myError.Message1 = result.Error.Error()
		myError.Message2 = utils.UpdateMysqlError
		return myError
	}
	return nil
}
