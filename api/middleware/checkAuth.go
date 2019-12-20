package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"newbug/util"
	"strings"
)

/**
 * 鉴权，校验token合法性
 */
func CheckUserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request

		//进行鉴权操作
		token := c.GetHeader("Authorization")
		if len(token) <= 0 {
			util.ResponseJson(c,http.StatusNotAcceptable,"token缺失",nil)
			c.Abort()
		}
		tokenSlice := strings.Split(token," ")
		token = tokenSlice[1]
		if len(token) <= 0 {
			util.ResponseJson(c,http.StatusNotAcceptable,"token缺失",nil)
			c.Abort()
		}

		var jwt util.Jwt
		jwt.Token = token
		parseErr := jwt.ParseToken()
		if parseErr != nil {
			util.Error(fmt.Sprintf("解析token失败：%v\n",parseErr.Error()))
			util.ResponseJson(c,http.StatusNotAcceptable,"非法token",nil)
			c.Abort()
		}

		if len(jwt.Email.(string)) <= 0 {
			util.Error(fmt.Sprintf("解析token失败：%v\n",token))
			util.ResponseJson(c,http.StatusNotAcceptable,"非法token",nil)
			c.Abort()
		}

		c.Set("email",jwt.Email)

		c.Next()

		// after request
	}
}
