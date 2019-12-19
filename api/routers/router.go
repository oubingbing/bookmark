package routers

import (
	"github.com/gin-gonic/gin"
	"newbug/controller"
	"newbug/controller/middleware"
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
		auth.GET("/register",controller.Register)
	}

	return router
}
