package utils

import (
	"github.com/gin-gonic/gin"
)

func ReturnError(c *gin.Context) {
	c.JSON(404, gin.H{
		"message": "error",
	})
}

func ReturnErrorString(c *gin.Context, errStr string) {
	c.JSON(404, gin.H{
		"message": errStr,
	})
}

func ReturnOk(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "succeed",
	})
}

func ReturnOkString(c *gin.Context, okStr string) {
	c.JSON(200, gin.H{
		"message": okStr,
	})
}
