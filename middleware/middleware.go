package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HiMiddlware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		fmt.Println("my customer middleware was called")
	}
}
