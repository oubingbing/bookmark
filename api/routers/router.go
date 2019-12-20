package routers

import (
	"github.com/gin-gonic/gin"
	"newbug/controller"
	"newbug/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	//需要鉴权后才能访问
	api := router.Group("/api")
	api.Use(middleware.CheckUserAuth())
	{
		api.GET("/user", controller.User)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/register",controller.Register)
		auth.POST("/login",controller.Login)
	}

	return router
}
