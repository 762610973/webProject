package router

import (
	"net/http"
	"webProject/controller"
	"webProject/logger"
	"webProject/middlewares"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置为发布模式，日志不会打印到控制台
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")
	// 注册
	v1.POST("/signUp", controller.SignUpHandler)
	// 登录
	v1.POST("/login", controller.LoginHandler)
	// 使用了中间件之后，之后所有的操作就都需要登陆了
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.POST("/setInfo", controller.SetInfoHandler)
		v1.GET("/posts", controller.GetPostListHandler)

		v1.GET("/post2", controller.GetPostListHandler2)

		v1.POST("/vote", controller.PostVoteHandler)
	}
	/*v1.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		// 如果是登录的用户,判断请求头中是否有有效的JWT
		c.Request.Header.Get("Authorization")
		if true {
			c.String(http.StatusOK, "pong")

		} else {
			// 否则直接返回请登录
			c.String(http.StatusOK, "请登录")

		}
	})*/
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
