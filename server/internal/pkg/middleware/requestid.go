package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/huangyul/miniblog/internal/pkg/konwn"
)

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		reqID := ctx.Request.Header.Get(konwn.XRequestIDKey)

		if reqID == "" {
			reqID = uuid.New().String()
		}

		ctx.Set(konwn.XRequestIDKey, reqID)

		ctx.Writer.Header().Set(konwn.XRequestIDKey, reqID)

		ctx.Next()
	}
}
