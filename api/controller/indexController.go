package controller

import "github.com/gin-gonic/gin"

func User(ctx *gin.Context)  {
	ctx.JSON(200,"user")
}
