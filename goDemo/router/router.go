package router

import (
	"goDemo/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/login", controller.Login)
	router.GET("/register", controller.Register)
	router.GET("/get", controller.GetUser)
	router.GET("/delete", controller.DeleteUser)
	router.GET("/update", controller.UpdateName)
	router.GET("/ping", pingHandler)
	router.GET("/ip", ipHandler)

	return router
}

// 测试用
func pingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})

}

func ipHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"ip": c.ClientIP(),
	})
}
