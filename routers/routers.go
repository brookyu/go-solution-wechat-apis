package routers

import (
	"github.com/brookyu/go-solution-wechat-apis/controllers/api"
	"github.com/brookyu/go-solution-wechat-apis/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//Giving access to storage folder
	r.Static("/storage", "storage")

	//Giving access to template folder
	r.Static("/templates", "templates")
	r.LoadHTMLGlob("templates/*")

	r.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	//API route for version 1
	publicGroup := r.Group("/api/public_group")

	//If you want to pass your route through specific middlewares
	publicGroup.Use(middlewares.UserMiddleWares())
	{
		publicGroup.POST("user-list", api.UserList)
	}

	//API route for version 2
	authenticatedGroup := r.Group("/api/auth").Use(middlewares.WechatAuthCheck())

	authenticatedGroup.POST("user-list", api.UserList)

	return r
}
