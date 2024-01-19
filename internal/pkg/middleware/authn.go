package middleware

import (
	"miniblog/internal/pkg/core"
	"miniblog/internal/pkg/errno"
	"miniblog/internal/pkg/known"
	"miniblog/pkg/token"

	"github.com/gin-gonic/gin"
)

func Authn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 解析 JWT token
		username, err := token.ParseRequest(ctx)
		if err != nil {
			core.WriteResponse(ctx, errno.ErrTokenInvaild, nil)
			ctx.Abort()

			return
		}

		ctx.Set(known.XUsernameKey, username)
		ctx.Next()
	}
}
