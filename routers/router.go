package routers

import (
	v1 "myapp/api/v1"
	"myapp/middleware"
	"myapp/utils"

	"github.com/gin-gonic/gin"
)

// 初始化路由
func InitRouter() {

	gin.SetMode(utils.AppMode)

	r := gin.New()
	r.Use()
	// todo
	// r.User(middleware.Logger())
	r.Use(gin.Recovery())

	//分组
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//文章路由模块
		auth.POST("article/add",v1.AddArt)
		auth.DELETE("article/:id",v1.DeleteArt)
	}

	router := r.Group("api/v1")
	{

		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.POST("login", v1.Login)
		router.GET("articles",v1.GetArt)
	}

	r.Run()
}
