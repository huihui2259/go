package controller

import (
	"encoding/json"
	"fmt"
	"goDemo/entity"
	"goDemo/service"
	"goDemo/utils"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func GetTest(c *gin.Context) {
	id := c.Param("id")
	s := service.GetShopTest1(id)
	utils.ReturnOkString(c, s)
}

func PostTest(c *gin.Context) {
	// json := make(map[string]interface{}) //注意该结构接受的内容
	// c.BindJSON(&json)
	// log.Printf("%v", &json)
	shop := &entity.Shop{}
	c.BindJSON(shop)
	log.Println(shop)
	service.SaveShop(shop)
	utils.ReturnOk(c)
}

func GetShopByID(c *gin.Context) {
	id := utils.AnalyzeID(c)
	if id == -1 {
		utils.ReturnErrorString(c, "id错误")
		return
	}
	shop, err := service.GetShopByID(id)
	if err != nil {
		utils.ReturnErrorString(c, err.String())
		return
	}
	shopJson, _ := json.Marshal(shop)
	utils.ReturnOkString(c, string(shopJson))
}

func GetShopListByType(c *gin.Context) {
	typeID, ok := c.GetQuery("type_id")
	if !ok {
		utils.ReturnErrorString(c, "typeID错误")
		return
	}
	TypeID, _ := strconv.Atoi(typeID)
	field := c.Query("field")
	shopList, err := service.GetShopListByTypeID(TypeID, field)
	if err != nil {
		utils.ReturnErrorString(c, err.String())
		return
	}
	listJson, _ := json.Marshal(shopList)
	utils.ReturnOkString(c, string(listJson))
}

func GetShopListByPage(c *gin.Context) {
	typeID, _ := strconv.Atoi(c.Query("type_id"))
	pageCount, _ := strconv.Atoi(c.Query("count"))
	pageIndex, _ := strconv.Atoi(c.Query("index"))
	field := c.Query("field")
	shopList, err := service.GetShopListByPage(typeID, pageIndex, pageCount, field)
	if err != nil {
		utils.ReturnErrorString(c, err.String())
		return
	}
	listJson, _ := json.Marshal(shopList)
	utils.ReturnOkString(c, string(listJson))
}

func UpdateShopByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	field := c.Query("field")
	value := c.Query("value")
	err := service.UpdateShopByID(id, field, value)
	if err != nil {
		utils.ReturnErrorString(c, err.String())
		return
	}
	utils.ReturnOk(c)
}

// 解决了缓存穿透的获取shop
func GetShopByIDChuanTou(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	shop, err := service.GetShopByIDChuanTou(id)
	if err != nil {
		utils.ReturnErrorString(c, err.String())
		return
	}
	jsons, _ := json.Marshal(shop)
	utils.ReturnOkString(c, string(jsons))
}

// 解决了缓存穿透和使用互斥锁缓存击穿的获取shop
func GetShopByIDJiChuan(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	shop, err := service.GetShopByIDJiChuan(id)
	if err != nil {
		utils.ReturnErrorString(c, err.String())
		return
	}
	jsons, _ := json.Marshal(shop)
	utils.ReturnOkString(c, string(jsons))
}

// 测试获取分布式锁
func GetLock(c *gin.Context) {
	key := c.Query("lock")
	value := uuid.New().String()
	flag := utils.TryLock(key, value)
	if flag {
		fmt.Println("获取锁成功,暂停20秒钟...")
		time.Sleep(20 * time.Second)
		utils.UnLock(key, value)
		fmt.Println("已解锁...")
		utils.ReturnOkString(c, "获取锁成功")
	}
}
