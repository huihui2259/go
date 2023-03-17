package controller

import (
	"encoding/json"
	"goDemo/mysql"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Save(c *gin.Context) {
	user := &mysql.AdminUser{}
	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(404, gin.H{
			"error": "未能找到ID",
		})
		return
	}
	user.ID, _ = strconv.Atoi(id)
	user.UserName = c.Query("name")
	user.Password = c.DefaultQuery("password", "123456")
	user.Tel = c.DefaultQuery("tel", "120")
	if err := mysql.SaveUser(user); err != nil {
		c.JSON(404, gin.H{
			"error": "存储玩家数据错误",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "succeed",
	})

}

func GetUser(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(404, gin.H{
			"error": "未能找到ID",
		})
	}
	ID, _ := strconv.Atoi(id)
	user, err := mysql.GetUser(ID)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "查询数据库出错",
		})
	}
	jsons, _ := json.Marshal(user)
	c.JSON(200, jsons)
}

func DeleteUser(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(404, gin.H{
			"error": "未能找到ID",
		})
		return
	}
	ID, _ := strconv.Atoi(id)
	if err := mysql.Delete(ID); err != nil {
		c.JSON(404, gin.H{
			"error": "删除数据库出错",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "succeed",
	})
}

func Update(c *gin.Context) {
	user := &mysql.AdminUser{}
	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(404, gin.H{
			"error": "未能找到ID",
		})
		return
	}
	user.ID, _ = strconv.Atoi(id)
	user.Addr = c.DefaultQuery("addr", "beijing")
	if err := mysql.Update(user); err != nil {
		c.JSON(404, gin.H{
			"error": "修改数据库出错",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "succeed",
	})
}
