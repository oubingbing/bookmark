package controller

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
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

func Test()  {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:%v/restricted", 8080), nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	res, err := http.DefaultClient.Do(req)
	fatal(err)

	// Read the response body
	buf := new(bytes.Buffer)
	io.Copy(buf, res.Body)
	res.Body.Close()
	fmt.Println(buf.String())
}
