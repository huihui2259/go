package utils

import (
	"goDemo/log2"
	"goDemo/redis"
	"io/ioutil"
)

// 为了防止出现以下情况：
// a获取到锁，但执行很久，锁过期，b获取到锁
// 此时a执行完，将b的锁释放，造成混乱，因此需要在设置锁的时候添加value
// 这里value有两种方案 1.使用当前协程的进程号或协程号 2.使用uuid，这里使用第二种方案

// 此外，该锁是不可重入的，若想实现可重入(同一个线程可以重复获取锁)，
// 可以存hash，key为线程(协程)ID，value为锁的次数，加锁时锁的次数加一
// 解锁时次数减一，若次数为0，就del
func TryLock(key, value string) bool {
	flag, err := redis.RedisClient.SetNX(key, value, LockExpireTime).Result()
	if err != nil {
		return false
	}
	return flag
}

func UnLock(key, value string) {
	file, _ := ioutil.ReadFile("/home/kingykwang/gitcode/go/goDemo/utils/lock.lua")
	// redis.RedisClient.Eval(string(file), []string{key}, value)
	result, _ := redis.RedisClient.Eval(string(file), []string{key}, value).Result()
	log2.Info.Printf("lua执行结果: %v", result)
	// redis.RedisClient.Del(key)
}
