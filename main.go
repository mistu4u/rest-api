package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api := initApi()
	r.GET("/hi", api.SayHi)
	r.Run()
}
