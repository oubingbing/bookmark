package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"newbug/util"
)

func Register(ctx *gin.Context)  {

	var jwt util.Jwt
	jwt.Email = "875307054@qq.com"

	err := jwt.CreateToken()
	if err != nil{
		fmt.Printf("创建token失败:%v\n",err.Error())
	}

	fmt.Println(jwt.Token)

	parseError := jwt.ParseToken()
	if parseError != nil{
		fmt.Printf("解析token失败:%v\n",parseError.Error())
	}

	ctx.JSON(200,jwt.Email)
}
