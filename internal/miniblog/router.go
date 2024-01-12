package miniblog

import (
	"miniblog/internal/pkg/core"
	"miniblog/internal/pkg/errno"
	"miniblog/internal/pkg/log"

	"github.com/gin-gonic/gin"
)

func installRouters(g *gin.Engine) error {
	g.NoRoute(func(ctx *gin.Context) {
		core.WriteResponse(ctx, errno.ErrPageNotFound, nil)
	})

	g.GET("/healthz", func(ctx *gin.Context) {
		log.C(ctx).Infow("Healthz function called")
	})

	return nil
}
