package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// NoCache 用来禁止客户端缓存 HTTP 请求返回的结果
func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))

	c.Next()
}

// Cors 用于设置 options 请求的返回头，设置跨域
func Cors(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

func Secure(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	// 浏览器不要嵌套当前页面，提高防御点击劫持攻击的能力
	c.Header("X-Frame-Options", "DENY")
	// 阻止浏览器从文件内容中嗅探出响应的 MIME 类型
	c.Header("X-Content-Type-Options", "nosniff")
	// 表示如果检测到XSS攻击，浏览器会阻止页面加载
	c.Header("X-XSS-Protection", "1; mode=block")
	// 它告诉浏览器通过HTTPS连接访问网站，并在接下来的31536000秒（一年）内强制使用HTTPS。这个设置有助于防止中间人攻击，确保用户连接到安全的服务器。
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}
}
