package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huangyul/miniblog/internal/miniblog/service"
	"github.com/huangyul/miniblog/internal/pkg/core"
)

type UserHandler struct {
	svc service.IUserService
}

func NewUserHandler(svc service.IUserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (u *UserHandler) Register(ctx *gin.Context) {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, core.Response{
			Code:    "10001",
			Message: "请求参数有误",
		})
		return
	}
	err := u.svc.Create(ctx, req.Email, req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, core.Response{
			Code:    "10001",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, core.Response{
		Code:    "0",
		Message: "创建成功",
	})

}
