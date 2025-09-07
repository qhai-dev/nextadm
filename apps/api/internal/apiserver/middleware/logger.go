package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("%s", c.HandlerName())
		//nlog.Info("info test", "method", c.Request.Method)
		c.Next()
	}
}
