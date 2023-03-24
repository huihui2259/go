package redis

import (
	"goDemo/utils"
	"log"
)

func GetUser(id string) (string, error) {
	result, err := RedisClient.Get(id).Result()
	return result, err
}

func SetUser(id, user string) error {
	if err := RedisClient.Set(utils.UserIDPrefix+id, user, 0).Err(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func DelUser(id string) error {
	if err := RedisClient.Del(utils.UserIDPrefix + id).Err(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
