package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"newbug/util"
)

func User(ctx *gin.Context)  {
	email := ctx.GetString("email")
	util.ResponseJson(ctx,http.StatusOK,"操作成功",email)
	return
}
