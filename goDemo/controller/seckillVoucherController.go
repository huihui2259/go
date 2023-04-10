package controller

import (
	"fmt"
	"goDemo/service"
	"goDemo/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 最简单的秒杀
func SecKillSimple(c *gin.Context) {
	voucherIDStr := c.Query("voucherID")
	userIDStr := c.Query("userID")
	voucherID, _ := strconv.Atoi(voucherIDStr)
	userID, _ := strconv.Atoi(userIDStr)
	err, orderID := service.SecKillSimple(voucherID, userID)
	res := fmt.Sprintf("orderID: %d", orderID)
	if err != nil {
		utils.ReturnErrorString(c, err.String()+res)
		return
	}
	utils.ReturnOkString(c, res)
}

// 乐观锁的秒杀
func SecKillOptimistic(c *gin.Context) {
	voucherIDStr := c.Query("voucherID")
	userIDStr := c.Query("userID")
	voucherID, _ := strconv.Atoi(voucherIDStr)
	userID, _ := strconv.Atoi(userIDStr)
	err, orderID := service.SecKillWithOptimistic(voucherID, userID)
	res := fmt.Sprintf("orderID: %d", orderID)
	if err != nil {
		utils.ReturnErrorString(c, err.String()+res)
		return
	}
	utils.ReturnOkString(c, res)
}

// 一人一单的秒杀
func SecKillSingleOrder(c *gin.Context) {
	voucherIDStr := c.Query("voucherID")
	userIDStr := c.Query("userID")
	voucherID, _ := strconv.Atoi(voucherIDStr)
	userID, _ := strconv.Atoi(userIDStr)
	err, orderID := service.SecKillWithSingle(voucherID, userID)
	res := fmt.Sprintf("orderID: %d", orderID)
	if err != nil {
		utils.ReturnErrorString(c, err.String()+res)
		return
	}
	utils.ReturnOkString(c, res)
}

// 一人一单加锁
func SecKillLockSingle(c *gin.Context) {
	voucherIDStr := c.Query("voucherID")
	userIDStr := c.Query("userID")
	voucherID, _ := strconv.Atoi(voucherIDStr)
	userID, _ := strconv.Atoi(userIDStr)
	err, orderID := service.SecKillWithLockSingle(voucherID, userID)
	res := fmt.Sprintf("orderID: %d", orderID)
	if err != nil {
		utils.ReturnErrorString(c, err.String()+res)
		return
	}
	utils.ReturnOkString(c, res)
}

func SecKillInit(c *gin.Context) {
	service.SecKillInit()
	utils.ReturnOk(c)
}
