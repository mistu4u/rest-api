package main

import (
	"rest-api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.HiMiddlware())

	api := initApi()
	r.GET("/hi", api.SayHi)
	r.Run()
}
