package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"miniblog/internal/pkg/known"
)

// RequestID 在请求 context 和 response 中注入 X-Request-ID
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(known.XRequestIDKey)

		if requestID == "" {
			requestID = uuid.New().String()
		}

		// 将 RequestID 放到 context
		c.Set(known.XRequestIDKey, requestID)

		// 将 RequestID 放到 response
		c.Writer.Header().Set(known.XRequestIDKey, requestID)

		c.Next()

	}
}
