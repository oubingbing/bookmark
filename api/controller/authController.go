package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"newbug/model"
	"newbug/service"
	"newbug/util"
)

func Register(ctx *gin.Context)  {

	var jwt util.Jwt
	jwt.Email = "875307054@qq.com"

	err := jwt.CreateToken()
	if err != nil{
		util.Error(fmt.Sprintf("创建token失败：%v\n",err.Error()))
	}

	fmt.Println(jwt.Token)

	parseError := jwt.ParseToken()
	if parseError != nil{
		util.Error(fmt.Sprintf("解析token失败：%v\n",parseError.Error()))
	}

	ctx.JSON(200,jwt.Email)
}

func Login(ctx *gin.Context)  {
	var user model.User
	service.FindByEmail(&user)

	resp := util.Response{200,"sdf","sdf"}
	ctx.JSON(http.StatusOK,resp)
}