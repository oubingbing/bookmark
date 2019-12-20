package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"newbug/model"
	"newbug/service"
	"newbug/util"
	"strings"
)

/**
 * 注册用户
 */
func Register(ctx *gin.Context)  {
	var registerUser model.RegisterUser
	if ctx.ShouldBind(&registerUser) != nil {
		util.ResponseJson(ctx,http.StatusInternalServerError,"参数错误",nil)
		return
	}

	validResult := service.Valid(registerUser)
	for _,v := range validResult {
		util.ResponseJson(ctx,http.StatusInternalServerError,v.ErrorMessage,nil)
		return
	}

	if strings.Compare(registerUser.Password,registerUser.ConfirmPassword) != 0 {
		util.ResponseJson(ctx,http.StatusInternalServerError,"两次输入密码不一致",nil)
		return
	}

	var user model.User
	user.Password = registerUser.Password
	user.Email 	  = registerUser.Email
	user.Nickname = registerUser.Nickname

	service.FindByEmail(&user)
	if user.ID > 0 {
		util.ResponseJson(ctx,http.StatusInternalServerError,"邮箱已存在",nil)
		return
	}

	user.Ip = ctx.ClientIP()
	user.Salt = rand.Intn(10000)
	user.Password = util.MD5Password(user.Password,user.Salt)
	result,RegisterErr := service.Store(&user)
	if result <= 0 || RegisterErr != nil {
		util.Error(fmt.Sprintf("注册失败：%v\n",RegisterErr.Error()))
		util.ResponseJson(ctx,http.StatusInternalServerError,"注册失败",nil)
		return
	}

	util.ResponseJson(ctx,http.StatusOK,"注册成功",nil)
	return
}

func Login(ctx *gin.Context)  {
	var user model.User
	if ctx.ShouldBind(&user) != nil {
		util.ResponseJson(ctx,http.StatusInternalServerError,"参数错误",nil)
		return
	}

	validResult := service.Valid(user)
	fmt.Println(validResult)
	for _,v := range validResult {
		util.ResponseJson(ctx,http.StatusInternalServerError,v.ErrorMessage,nil)
		return
	}

	password := user.Password
	service.FindByEmail(&user)
	if user.ID <= 0 {
		util.ResponseJson(ctx,http.StatusInternalServerError,"账号或密码错误",nil)
		return
	}

	if strings.Compare(user.Password,util.MD5Password(password,user.Salt)) != 0 {
		util.ResponseJson(ctx,http.StatusInternalServerError,"账号或密码错误",nil)
		return
	}

	//发放token
	var jwt util.Jwt
	jwt.Email = user.Email
	err := jwt.CreateToken()
	if err != nil{
		util.Error(fmt.Sprintf("创建token失败：%v\n",err.Error()))
	}

	util.ResponseJson(ctx,http.StatusInternalServerError,"登录成功",jwt.Token)
	return
}