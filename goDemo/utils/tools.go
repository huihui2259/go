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
