package service

import (
	"encoding/json"
	"fmt"
	"goDemo/entity"
	"goDemo/log2"
	"goDemo/mysql"
	"goDemo/redis"
	"goDemo/utils"
	"log"
	"strconv"
	"time"
)

func GetUser() *entity.User {
	// 模拟登录,每次登录不同id
	now := time.Now().Unix()

	id := int(now%6) + 1
	user, _ := GetUserByID(id)
	log2.Info.Printf("当前登录的用户 %v", user)
	return user
}

func GetUserNow() *entity.User {
	// 模拟登录,每次只登录id=1
	id := 1
	user, _ := GetUserByID(id)
	log2.Info.Printf("当前登录的用户 %v", user)
	return user
}

// func GetUsers() *[]entity.User {
// 	// 模拟登录
// 	id := 1
// 	user, _ := GetUserByID(id)
// 	return user
// }

func GetUsersByIDs(ids []int) *[]entity.User {
	users := &[]entity.User{}
	query := fmt.Sprintf("select * from tb_user where id in (%s) order by field(id,%s)", utils.StringJoin(ids, ","), utils.StringJoin(ids, ","))
	log2.Info.Println(query)
	mysql.Db.Model(&entity.User{}).Raw(query).Find(users)
	log2.Info.Println(users)
	return users
	// mysql.Db.Model(&entity.User{}).Where("id IN ?", ids).Order("FIELD(id,?)", ids).Find(users)
}

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
	if err := redis.RedisClient.Set(utils.UserIDPrefix+strconv.Itoa(id), string(jsons), utils.CommonExpireTime).Err(); err != nil {
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
	if err := redis.RedisClient.Set(utils.UserIDPrefix+strconv.Itoa(user.ID), string(jsons), utils.CommonExpireTime).Err(); err != nil {
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
