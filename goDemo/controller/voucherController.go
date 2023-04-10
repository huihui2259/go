package controller

import (
	"encoding/json"
	"goDemo/entity"
	"goDemo/service"
	"goDemo/utils"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取优惠券信息(传入优惠券id或商店id)
func GetVoucher(c *gin.Context) {
	para, ok := c.GetQuery("shop_id")
	if ok {
		id, _ := strconv.Atoi(para)
		voucherList, err := service.GetVoucherListByShopID(id)
		if err != nil {
			utils.ReturnErrorString(c, err.String())
			return
		}
		jsons, _ := json.Marshal(voucherList)
		utils.ReturnOkString(c, string(jsons))
		return
	}
	para, ok = c.GetQuery("id")
	if ok {
		id, _ := strconv.Atoi(para)
		voucher, err := service.GetVoucherByID(id)
		if err != nil {
			utils.ReturnErrorString(c, err.String())
			return
		}
		jsons, _ := json.Marshal(voucher)
		utils.ReturnOkString(c, string(jsons))
		return
	}
}

func AddSecKillVoucher(c *gin.Context) {
	// 增加秒杀券，秒杀券首先是优惠券
	voucher := &entity.Voucher{}
	c.BindJSON(voucher)
	log.Println(voucher)
	err := service.AddSeckillVoucher(voucher)
	if err != nil {
		utils.ReturnErrorString(c, err.String())
		return
	}
	utils.ReturnOk(c)
}
