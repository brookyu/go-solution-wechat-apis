package main

import (
	"fmt"
	_ "github.com/brookyu/go-solution-wechat-apis/docs"
	"github.com/brookyu/go-solution-wechat-apis/routers"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	r := routers.SetupRouter()
	port := os.Getenv("port")
	url := ginSwagger.URL("http://localhost:8099/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	if len(os.Args) > 1 {
		reqPort := os.Args[1]
		if reqPort != "" {
			port = reqPort
		}
	}
	r.Run(":" + port)
}
