package service

import (
	"goDemo/log2"
	"goDemo/redis"
	"goDemo/utils"
	"strconv"
	"time"
)

// 存储玩家签到数据
func UserSign(id int) *utils.MyError {
	myErr := &utils.MyError{}
	key := SignGetKey(id)
	days := time.Now().Format("02")
	offset, _ := strconv.Atoi(days)
	err := redis.RedisClient.SetBit(key, int64(offset-1), utils.TRUE).Err()
	if err != nil {
		myErr.Message1 = err.Error()
		myErr.Message2 = utils.SetRedisError
		return myErr
	}
	return nil
}

// 获取玩家某天是否已经签到
func GetIsSigned(id int, stime string) (*utils.MyError, bool) {
	myErr := &utils.MyError{}
	theTime := utils.StringToTime(stime, utils.TimeFormat1)
	dayInt, _ := strconv.Atoi(theTime.Format(utils.Day))

	offset := dayInt - 1
	key := utils.SignPrefix + theTime.Format(utils.Year) + theTime.Format(utils.Month) + ":" + strconv.Itoa(id)

	log2.Info.Printf("offset: %d", offset)
	log2.Info.Printf("key: %s", key)
	isSign, err := redis.RedisClient.GetBit(key, int64(offset)).Result()
	if err != nil {
		myErr.Message1 = err.Error()
		myErr.Message2 = utils.SetRedisError
		return myErr, false
	}
	if isSign != 0 {
		return nil, true
	}
	return nil, false
}

// 获取连续签到天数
func SumContinueSign(id int) (*utils.MyError, int) {
	myErr := &utils.MyError{}
	key := SignGetKey(id)
	days := time.Now().Format(utils.Day)
	offset, _ := strconv.Atoi(days)
	offset = offset - 1
	count := 0
	for ; offset >= 0; offset = offset - 1 {
		isSign, err := redis.RedisClient.GetBit(key, int64(offset)).Result()
		if err != nil {
			myErr.Message1 = err.Error()
			myErr.Message2 = utils.SetRedisError
			return myErr, 0
		}
		if isSign == 0 {
			break
		}
		count++
	}
	return nil, count
}

func RemedySign(id int, stime string) (*utils.MyError, bool) {
	myErr := &utils.MyError{}
	myErr1, isSign := GetIsSigned(id, stime)
	if myErr1 != nil || isSign {
		return myErr, false
	}
	theTime := utils.StringToTime(stime, utils.TimeFormat1)
	dayInt, _ := strconv.Atoi(theTime.Format(utils.Day))

	offset := dayInt - 1
	key := utils.SignPrefix + theTime.Format(utils.Year) + theTime.Format(utils.Month) + ":" + strconv.Itoa(id)

	err := redis.RedisClient.SetBit(key, int64(offset), utils.TRUE).Err()
	if err != nil {
		myErr.Message1 = err.Error()
		myErr.Message2 = utils.SetRedisError
		return myErr, false
	}
	return nil, true
}

func SignGetKey(id int) string {
	now := time.Now()
	year := now.Format("2006")
	mouth := now.Format("01")
	idStr := strconv.Itoa(id)
	return utils.SignPrefix + year + mouth + ":" + idStr
}
