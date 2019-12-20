package middleware

import (
	"github.com/gin-gonic/gin"
)

func Format() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request

		c.Next()

		c.GetHeader("Authorization")

		// after request
	}
}