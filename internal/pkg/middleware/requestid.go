package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"miniblog/internal/pkg/know"
)

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := ctx.Request.Header.Get(know.XRequestIDKey)

		if requestID == "" {
			requestID = uuid.NewString()
		}

		// 保存到上下文
		ctx.Set(know.XRequestIDKey, requestID)

		// 保存到 HTTP 返回中
		ctx.Writer.Header().Set(know.XRequestIDKey, requestID)

		ctx.Next()
	}
}
