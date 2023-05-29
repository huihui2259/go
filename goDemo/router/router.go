package router

import (
	"goDemo/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	userRouter := router.Group("/user")
	{
		userRouter.GET("/login", controller.Login)
		userRouter.GET("/register", controller.Register)
		userRouter.GET("/get", controller.GetUser)
		userRouter.GET("/delete", controller.DeleteUser)
		userRouter.GET("/update", controller.UpdateName)
		// http://9.135.34.52:8089/user/save
		userRouter.POST("/save", controller.SaveTest)
	}
	shopRouter := router.Group("/shop")
	{
		shopRouter.GET("/test/:id", controller.GetTest)
		shopRouter.POST("/post", controller.PostTest)
		shopRouter.GET("/get", controller.GetShopByID)
		shopRouter.GET("/getType", controller.GetShopListByType)
		// http://9.135.34.52:8089/shop/getPage?type_id=1&index=1&count=3&field=*
		shopRouter.GET("/getPage", controller.GetShopListByPage)
		// http://9.135.34.52:8089/shop/update?id=1&index=3&field=area&value=中关村
		shopRouter.GET("/update", controller.UpdateShopByID)
		// http://9.135.34.52:8089/shop/chuantou?id=1
		shopRouter.GET("/chuantou", controller.GetShopByIDChuanTou)
		// http://9.135.34.52:8089/shop/lock?lock=lock
		shopRouter.GET("/lock", controller.GetLock)
		// http://9.135.34.52:8089/shop/jichuan?id=1
		shopRouter.GET("/jichuan", controller.GetShopByIDJiChuan)
	}

	voucherRouter := router.Group("/voucher")
	{
		// http://9.135.34.52:8089/voucher/get?id=1
		voucherRouter.GET("/get", controller.GetVoucher)
		// http://9.135.34.52:8089/voucher/add
		voucherRouter.POST("/add", controller.AddSecKillVoucher)

	}
	seckillRouter := router.Group("/seckill")
	{
		// http://9.135.34.52:8089/seckill/simple
		seckillRouter.GET("/simple", controller.SecKillSimple)
		// http://9.135.34.52:8089/seckill/optimistic
		seckillRouter.GET("/optimistic", controller.SecKillOptimistic)
		// http://9.135.34.52:8089/seckill/init
		seckillRouter.GET("/init", controller.SecKillInit)
		// http://9.135.34.52:8089/seckill/single
		seckillRouter.GET("/single", controller.SecKillSingleOrder)
		// http://9.135.34.52:8089/seckill/lockSingle
		seckillRouter.GET("/lockSingle", controller.SecKillLockSingle)
		// http://9.135.34.52:8089/seckill/kafkaV1
		seckillRouter.GET("/kafkaV1", controller.SecKillKafkaV1)
		// http://9.135.34.52:8089/seckill/final
		seckillRouter.GET("/final", controller.SecKillFinal)
	}

	blogRouter := router.Group("/blog")
	{
		blogRouter.GET("/hot", controller.GetHotBlog)
		// http://9.135.34.52:8089/blog/5
		blogRouter.GET("/:id", controller.GetBlogByID)
		// http://9.135.34.52:8089/blog/like/5
		blogRouter.GET("/like/:blogID", controller.LikeBlog)
		// http://9.135.34.52:8089/blog/queryLikes/5
		blogRouter.GET("/queryLikes/:blogID", controller.QueryLikes)
		// http://9.135.34.52:8089/blog/followV1/2
		blogRouter.GET("/followV1/:userID", controller.FollowV1)
		// http://9.135.34.52:8089/blog/followV2/4/2
		blogRouter.GET("/followV2/:myID/:userID", controller.FollowV2)
		// http://9.135.34.52:8089/blog/common/4/1
		blogRouter.GET("/common/:myID/:userID", controller.CommonFollow)
	}

	signRouter := router.Group("/sign")
	{
		// http://9.135.34.52:8089/sign/1
		signRouter.GET("/:id", controller.Sign)
		// http://9.135.34.52:8089/sign/1/20230523
		signRouter.GET("/:id/:date", controller.IsSign)
		// http://9.135.34.52:8089/sign/sum/1
		signRouter.GET("/sum/:id", controller.SumSign)
		// http://9.135.34.52:8089/sign/remedy/1/20230523
		signRouter.GET("/remedy/:id/:date", controller.RemedySign)
	}

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
