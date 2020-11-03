package main

import (
	"fmt"
	"github.com/brookyu/go-solution-wechat-apis/routers"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	r := routers.SetupRouter()
	port := os.Getenv("port")
	if len(os.Args) > 1 {
		reqPort := os.Args[1]
		if reqPort != "" {
			port = reqPort
		}
	}
	r.Run(":" + port)
}
