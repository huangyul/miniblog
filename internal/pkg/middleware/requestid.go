package middleware

import (
	"miniblog/internal/pkg/known"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestID := ctx.Request.Header.Get(known.XRequestIDKey)

		if requestID == "" {
			requestID = uuid.NewString()
		}

		// 保存一份到 ctx 中
		ctx.Set(known.XRequestIDKey, requestID)

		ctx.Writer.Header().Set(known.XRequestIDKey, requestID)
		ctx.Next()
	}
}
