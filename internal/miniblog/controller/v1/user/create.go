package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"miniblog/internal/pkg/core"
	"miniblog/internal/pkg/errno"
	"miniblog/internal/pkg/log"
	v1 "miniblog/pkg/api/miniblog/v1"
)

func (c *UserController) Create(ctx *gin.Context) {
	log.C(ctx).Infow("create user function called")

	var r v1.CreateUserRequest
	if err := ctx.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(ctx, errno.ErrBind, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(ctx, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	if err := c.b.Users().Create(ctx, &r); err != nil {
		core.WriteResponse(ctx, err, nil)

		return
	}

	core.WriteResponse(ctx, nil, nil)
}
