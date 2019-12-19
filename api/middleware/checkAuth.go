package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CheckUserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		fmt.Println("请求前")

		c.Next()

		c.GetHeader("Authorization")

		// after request
		fmt.Printf("请求后")
	}
}
