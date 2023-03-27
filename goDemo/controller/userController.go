package controller

import (
	"encoding/json"
	"goDemo/entity"
	"goDemo/service"
	"goDemo/utils"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	user, err := service.GetUserByID(ID)
	if err != nil {
		utils.ReturnErrorString(c, "请先注册")
	}
	if user.Password != password {
		utils.ReturnErrorString(c, "密码错误")
	}
	utils.ReturnOkString(c, "登录成功")
}

func Register(c *gin.Context) {
	user := &entity.User{}
	id := c.Query("id")
	user.ID, _ = strconv.Atoi(id)
	user.NickName = c.Query("nick_name")
	user.Password = c.DefaultQuery("password", "123456")
	user.Phone = c.DefaultQuery("phone", "120")
	user.CreateTime = time.Now().Format(utils.TimeFormat)
	user.UpdateTime = time.Now().Format(utils.TimeFormat)
	log.Println(user)
	if err := service.RegisterUser(user); err != nil {
		utils.ReturnErrorString(c, err.String())
		return
	}
	utils.ReturnOkString(c, "注册成功")
}

func GetUser(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		utils.ReturnErrorString(c, "请输入ID")
		return
	}
	ID, _ := strconv.Atoi(id)
	user, err := service.GetUserByID(ID)
	if err != nil {
		utils.ReturnErrorString(c, err.String())
	}
	// 返回给前端
	jsons, _ := json.Marshal(user)
	utils.ReturnOkString(c, string(jsons))
}

func DeleteUser(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		utils.ReturnErrorString(c, "请输入ID")
		return
	}
	ID, _ := strconv.Atoi(id)
	if err := service.Delete(ID); err != nil {
		utils.ReturnErrorString(c, err.String())
		return
	}
	utils.ReturnOkString(c, "注销成功")
}

// 修改昵称
func UpdateName(c *gin.Context) {
	user := &entity.User{}
	id, ok := c.GetQuery("id")
	if !ok {
		utils.ReturnErrorString(c, "请输入ID")
		return
	}
	user.ID, _ = strconv.Atoi(id)
	user.NickName = c.DefaultQuery("nick_name", uuid.New().String()[0:10])
	user.UpdateTime = time.Now().Format(utils.TimeFormat)
	if err := service.UpdateNickName(user); err != nil {
		utils.ReturnErrorString(c, err.String())
		return
	}
	utils.ReturnOkString(c, "改名成功")
}
