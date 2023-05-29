package controller

import (
	"goDemo/service"
	"goDemo/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Sign(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	err := service.UserSign(id)
	if err != nil {
		utils.ReturnErrorString(c, err.String())
		return
	}
	utils.ReturnOk(c)
}

func IsSign(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	timeStr := c.Param("date")
	err, isSign := service.GetIsSigned(id, timeStr)
	if err != nil {
		utils.ReturnErrorString(c, err.String())
		return
	}
	if isSign {
		utils.ReturnOkString(c, "1")
		return
	}
	utils.ReturnOkString(c, "0")
}

func SumSign(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	err, sum := service.SumContinueSign(id)
	if err != nil {
		utils.ReturnErrorString(c, err.String())
		return
	}
	utils.ReturnOkString(c, strconv.Itoa(sum))
}

func RemedySign(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	timeStr := c.Param("date")
	err, isSucceed := service.RemedySign(id, timeStr)
	if err != nil {
		utils.ReturnErrorString(c, err.String())
		return
	}
	if isSucceed {
		utils.ReturnOkString(c, "1")
		return
	}
	utils.ReturnOkString(c, "0")
}
