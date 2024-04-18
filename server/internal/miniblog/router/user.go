package router

import (
	"github.com/gin-gonic/gin"
	"github.com/huangyul/miniblog/internal/miniblog/handler"
)

type UserRouter struct{}

func (u *UserRouter) RegisterUserRoute(r *gin.Engine, handler handler.UserHandler) {
	rg := r.Group("user")
	{
		rg.POST("register", handler.Register)
	}
}
