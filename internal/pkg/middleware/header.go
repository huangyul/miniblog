package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	})
}
