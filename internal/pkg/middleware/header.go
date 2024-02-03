package middleware

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 设置允许跨域
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	})
}

// NoCache 设置无缓存
func NoCache() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
		ctx.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
		ctx.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
		ctx.Next()
	}
}
