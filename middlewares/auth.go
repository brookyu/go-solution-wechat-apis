package middlewares

import "github.com/gin-gonic/gin"

func UserMiddleWares() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
