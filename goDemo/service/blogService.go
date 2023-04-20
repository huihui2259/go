package service

import (
	"goDemo/entity"
	"goDemo/log2"
	"goDemo/mysql"
	"goDemo/redis"
	"goDemo/utils"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"

	Redis "github.com/go-redis/redis"
)

// 查询点赞最高的几篇
func GetHotBlog(page int) (*[]entity.Blog, *utils.MyError) {
	blogList := &[]entity.Blog{}
	myError := &utils.MyError{}
	result := mysql.Db.Model(&entity.Blog{}).Order("liked DESC").Limit(utils.PageSize).Offset((page - 1) * utils.PageSize).Find(blogList)
	if result.RowsAffected == 0 {
		myError.Message2 = utils.GetEmptyDataError
		return blogList, myError
	}
	if result.Error != nil {
		myError.Message1 = result.Error.Error()
		myError.Message2 = utils.GetMysqlError
		return blogList, myError
	}
	return blogList, nil
}

// 根据id获取笔记
func GetBlogByID(id int) (*entity.Blog, *utils.MyError) {
	blog := &entity.Blog{}
	myError := &utils.MyError{}
	result := mysql.Db.Model(&entity.Blog{}).Where("id = ?", id).Find(blog)
	if result.RowsAffected == 0 {
		myError.Message2 = utils.GetEmptyDataError
		return blog, myError
	}
	if result.Error != nil {
		myError.Message1 = result.Error.Error()
		myError.Message2 = utils.GetMysqlError
		return blog, myError
	}
	return blog, nil
}

// 对某笔记点赞(第一次点击点赞，第二次点击取消点赞)
func LikeBlog(blogID int) *utils.MyError {
	user := GetUser()
	myError := &utils.MyError{}
	key := utils.BlogLikedIDPrefix + strconv.Itoa(blogID)
	score, err := redis.RedisClient.ZScore(key, strconv.Itoa(user.ID)).Result()
	if err != nil && err != Redis.Nil {
		// 查询出错
		myError.Message1 = err.Error()
		myError.Message2 = utils.ValueNullError
		return myError
	}
	if err == Redis.Nil {
		// 结果未找到，就点赞
		z := Redis.Z{Score: float64(time.Now().Unix()), Member: strconv.Itoa(user.ID)}
		redis.RedisClient.ZAdd(key, z)
		mysql.Db.Model(&entity.Blog{}).Update("liked", gorm.Expr("liked + 1")).Where("id = ?", blogID)
		log2.Info.Println("进行点赞")
	} else {
		// 已点过赞,就取消点赞
		log2.Info.Printf("score: %f\n", score)
		mysql.Db.Model(&entity.Blog{}).Update("liked", gorm.Expr("liked - 1")).Where("id = ?", blogID)
		redis.RedisClient.ZRem(key, strconv.Itoa(user.ID))
		log2.Info.Println("取消点赞")
	}
	return nil
}

// 某笔记的前五位点赞的人，用于展示
// 也可以认为是排行榜
func QueryLikes(blogID int) (*[]entity.User, *utils.MyError) {
	userList := &[]entity.User{}
	//myError := &utils.MyError{}
	key := utils.BlogLikedIDPrefix + strconv.Itoa(blogID)
	ids, _ := redis.RedisClient.ZRange(key, utils.Zero, 5).Result()
	log2.Info.Println(ids)
	idsInt := utils.StringSliceToInt(ids)
	log2.Info.Println(idsInt)
	userList = GetUsersByIDs(idsInt)
	return userList, nil
}

// 当前登录用户是否关注userID用户
func IsFollow(userID int) bool {
	userNow := GetUserNow()
	count := 0
	mysql.Db.Model(&entity.Follow{}).Where("user_id = ?", userNow.ID).Where("follow_user_id = ?", userID).Count(&count)
	return count != 0
}

// 指定用户是否关注userID用户
func IsFollowV2(userID, myID int) bool {
	count := 0
	mysql.Db.Model(&entity.Follow{}).Where("user_id = ?", myID).Where("follow_user_id = ?", userID).Count(&count)
	return count != 0
}

// 当前用户点击关注按钮userID用户
func Follow(userID int) (*utils.MyError, bool) {
	userNow := GetUserNow()
	myError := &utils.MyError{}
	if userNow.ID == userID {
		myError.Message2 = "不能关注自己"
		return myError, false
	}
	key := utils.FollowIDPrefix + strconv.Itoa(userNow.ID)
	if IsFollow(userID) {
		// 已关注，就取消关注
		result := mysql.Db.Model(&entity.Follow{}).Where("user_id = ? and follow_user_id = ?", userNow.ID, userID).Delete(&entity.Follow{})
		if result.RowsAffected != 0 {
			// 删除成功
			redis.RedisClient.SRem(key, strconv.Itoa(userID))
			return nil, false
		} else {
			// 数据库出错
			myError.Message1 = result.Error.Error()
			myError.Message2 = utils.DeleteMysqlError
			return myError, false
		}
	} else {
		// 未关注，就关注
		createTime := utils.GetNowTimeString()
		item := &entity.Follow{
			UserID:       userNow.ID,
			FollowUserID: userID,
			CreateTime:   createTime,
		}
		result := mysql.Db.Model(&entity.Follow{}).Create(item)
		if result.RowsAffected != 0 {
			// 更新成功，写入redis
			log2.Info.Printf("RowsAffected: %d", result.RowsAffected)
			redis.RedisClient.SAdd(key, utils.ToString(userID))
			return nil, true
		} else {
			// 数据库出错
			myError.Message1 = result.Error.Error()
			myError.Message2 = utils.InsertMysqlError
			return myError, false
		}
	}
}

// 指定用户点击关注按钮userID用户
func FollowV2(userID, myID int) (*utils.MyError, bool) {
	userNow, _ := GetUserByID(myID)
	flag := false
	myError := &utils.MyError{}
	if myID == userID {
		myError.Message2 = "不能关注自己"
		return myError, false
	}
	key := utils.FollowIDPrefix + strconv.Itoa(myID)
	if IsFollowV2(userID, userNow.ID) {
		// 已关注，就取消关注
		result := mysql.Db.Model(&entity.Follow{}).Where("user_id = ? and follow_user_id = ?", userNow.ID, userID).Delete(&entity.Follow{})
		if result.RowsAffected != 0 {
			// 删除成功
			redis.RedisClient.SRem(key, strconv.Itoa(userID))
			return nil, flag
		} else {
			// 数据库出错
			myError.Message1 = result.Error.Error()
			myError.Message2 = utils.DeleteMysqlError
			return myError, flag
		}
	} else {
		// 未关注，就关注
		createTime := utils.GetNowTimeString()
		item := &entity.Follow{
			UserID:       userNow.ID,
			FollowUserID: userID,
			CreateTime:   createTime,
		}
		result := mysql.Db.Model(&entity.Follow{}).Create(item)
		if result.RowsAffected != 0 {
			// 更新成功，写入redis
			log2.Info.Printf("RowsAffected: %d", result.RowsAffected)
			redis.RedisClient.SAdd(key, utils.ToString(userID))
			return nil, true
		} else {
			// 数据库出错
			myError.Message1 = result.Error.Error()
			myError.Message2 = utils.InsertMysqlError
			return myError, flag
		}
	}
}

func CommonFollow(userID, myID int) (*[]entity.User, *utils.MyError) {
	users := &[]entity.User{}
	myError := &utils.MyError{}
	key1 := utils.FollowIDPrefix + utils.ToString(userID)
	key2 := utils.FollowIDPrefix + utils.ToString(myID)
	result, err := redis.RedisClient.SInter(key1, key2).Result()
	if err != nil {
		myError.Message1 = err.Error()
		myError.Message2 = utils.GetRedisError
		return users, myError
	}
	log2.Info.Println(result)
	ids := utils.StringSliceToInt(result)
	users = GetUsersByIDs(ids)
	return users, nil
}
