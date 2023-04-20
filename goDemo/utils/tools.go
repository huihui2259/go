package utils

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AnalyzeID(c *gin.Context) int {
	id, ok := c.GetQuery("id")
	if !ok {
		return -1
	}
	ID, err := strconv.Atoi(id)

	if err != nil {
		return -1
	}
	return ID
}

func GetNowTimeString() string {
	now := time.Now()
	return now.Format(TimeFormat)
}

func StringToUnix(timeStr string) int64 {
	loc, _ := time.LoadLocation("Local")                         //获取时区
	theTime, _ := time.ParseInLocation(TimeFormat, timeStr, loc) //使用模板在对应时区转化为time.time类型
	return theTime.Unix()
}

func UnixToString(timeUnix int64) string {
	timeNow := time.Unix(timeUnix, 0)
	return timeNow.Format(TimeFormat)
}

func StringJoin(slice []int, s string) string {
	var res string
	for k, value := range slice {
		res += strconv.Itoa(value)
		if k != len(slice)-1 {
			res += s
		}
	}
	return res
}

func StringSliceToInt(slice []string) []int {
	res := []int{}
	for _, value := range slice {
		valueToInt, _ := strconv.Atoi(value)
		res = append(res, valueToInt)
	}
	return res
}

func ToString(value int) string {
	s := strconv.Itoa(value)
	return s
}

func ToInt(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}
