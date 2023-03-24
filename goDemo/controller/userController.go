package controller

import (
	"encoding/json"
	"goDemo/mysql"
	"goDemo/redis"
	"goDemo/utils"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		utils.ReturnErrorString(c, "请输入ID")
		return
	}
	password, ok := c.GetQuery("password")
	if !ok {
		utils.ReturnErrorString(c, "请输入密码")
		return
	}
	ID, _ := strconv.Atoi(id)
	user, err := mysql.GetUser(ID)
	if err != nil {
		utils.ReturnErrorString(c, "请先注册")
	}
	if user.Password != password {
		utils.ReturnErrorString(c, "密码错误")
	}
	utils.ReturnOkString(c, "登录成功")
}

func Save(c *gin.Context) {
	user := &mysql.User{}
	id := c.Query("id")
	user.ID, _ = strconv.Atoi(id)
	user.NickName = c.Query("nick_name")
	user.Password = c.DefaultQuery("password", "123456")
	user.Phone = c.DefaultQuery("phone", "120")
	user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	user.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	log.Println(user)
	if err := mysql.SaveUser(user); err != nil {
		utils.ReturnErrorString(c, "存储玩家数据错误")
		return
	}
	utils.ReturnOk(c)

}

func GetUser(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		utils.ReturnErrorString(c, "请输入ID")
		return
	}
	// 首先从redis中查询缓存
	if result, err := redis.GetUser(id + utils.UserIDPrefix); err == nil {
		// redis能找到
		utils.ReturnOkString(c, result)
		return
	}
	// 缓存找不到，就找数据库
	ID, _ := strconv.Atoi(id)
	user, err := mysql.GetUser(ID)
	if err != nil {
		utils.ReturnErrorString(c, "查询数据库出错")
		return
	}
	// 数据库找到，写回缓存
	jsons, _ := json.Marshal(user)
	if err := redis.SetUser(id, string(jsons)); err != nil {
		utils.ReturnError(c)
		return
	}
	utils.ReturnOkString(c, string(jsons))
}

func DeleteUser(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		utils.ReturnErrorString(c, "请输入ID")
		return
	}
	ID, _ := strconv.Atoi(id)
	if err := mysql.Delete(ID); err != nil {
		utils.ReturnErrorString(c, "删除数据库出错")
		return
	}
	utils.ReturnOk(c)
}

func Update(c *gin.Context) {
	user := &mysql.User{}
	id, ok := c.GetQuery("id")
	if !ok {
		utils.ReturnErrorString(c, "请输入ID")
		return
	}
	user.ID, _ = strconv.Atoi(id)
	user.UpdateTime = time.Now().Format("1998-05-04 00:00:00")
	if err := mysql.Update(user); err != nil {
		c.JSON(404, gin.H{
			"error": "修改数据库出错",
		})
		return
	}
	utils.ReturnOk(c)
}
