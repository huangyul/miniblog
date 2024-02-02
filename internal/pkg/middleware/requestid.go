package middleware

import (
	know "miniblog/internal/pkg/konw"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := ctx.Request.Header.Get(know.XRequestID)

		if requestId == "" {
			requestId = uuid.NewString()
		}

		ctx.Set(know.XRequestID, requestId)

		ctx.Header(know.XRequestID, requestId)

		ctx.Next()

	}
}
