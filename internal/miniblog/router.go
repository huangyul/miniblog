package miniblog

import (
	"miniblog/internal/miniblog/controller/v1/user"
	"miniblog/internal/miniblog/store"
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
		log.C(ctx).Infow("healthz function called")

		core.WriteResponse(ctx, nil, map[string]string{"status": "ok"})
	})

	uc := user.New(store.S)

	v1 := g.Group("/v1")
	{
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)
		}
	}

	return nil
}
