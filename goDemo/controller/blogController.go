package controller

import (
	"encoding/json"
	"goDemo/service"
	"goDemo/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetHotBlog(c *gin.Context) {
	page := c.Query("page")
	pageID, _ := strconv.Atoi(page)
	blogList, err := service.GetHotBlog(pageID)
	if err != nil {
		utils.ReturnErrorString(c, err.String())
		return
	}
	jsons, _ := json.Marshal(blogList)
	utils.ReturnOkString(c, string(jsons))
}

func GetBlogByID(c *gin.Context) {
	id := c.Param("id")
	ID, _ := strconv.Atoi(id)
	blog, err := service.GetBlogByID(ID)
	if err != nil {
		utils.ReturnErrorString(c, err.String())
		return
	}
	jsons, _ := json.Marshal(blog)
	utils.ReturnOkString(c, string(jsons))
}

func LikeBlog(c *gin.Context) {
	id := c.Param("blogID")
	blogID, _ := strconv.Atoi(id)
	err := service.LikeBlog(blogID)
	if err != nil {
		utils.ReturnErrorString(c, err.String())
		return
	}
	utils.ReturnOk(c)
}

func QueryLikes(c *gin.Context) {
	id := c.Param("blogID")
	blogID, _ := strconv.Atoi(id)
	userList, _ := service.QueryLikes(blogID)
	jsons, _ := json.Marshal(userList)
	utils.ReturnOkString(c, string(jsons))
}

func FollowV1(c *gin.Context) {
	id := c.Param("userID")
	userID := utils.ToInt(id)
	myError, flag := service.Follow(userID)
	if myError != nil {
		utils.ReturnErrorString(c, myError.String())
		return
	} else {
		if flag {
			utils.ReturnOkString(c, "关注成功")
			return
		} else {
			utils.ReturnOkString(c, "取关成功")
			return
		}
	}
}

func FollowV2(c *gin.Context) {
	myID := utils.ToInt(c.Param("myID"))
	userID := utils.ToInt(c.Param("userID"))
	myError, flag := service.FollowV2(userID, myID)
	if myError != nil {
		utils.ReturnErrorString(c, myError.String())
		return
	} else {
		if flag {
			utils.ReturnOkString(c, "关注成功")
			return
		} else {
			utils.ReturnOkString(c, "取关成功")
			return
		}
	}
}

func CommonFollow(c *gin.Context) {
	myID := utils.ToInt(c.Param("myID"))
	userID := utils.ToInt(c.Param("userID"))
	users, err := service.CommonFollow(userID, myID)
	if err != nil {
		utils.ReturnErrorString(c, err.String())
		return
	}
	jsons, _ := json.Marshal(users)
	utils.ReturnOkString(c, string(jsons))
}
