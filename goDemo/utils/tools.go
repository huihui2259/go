package utils

import (
	"strconv"

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
