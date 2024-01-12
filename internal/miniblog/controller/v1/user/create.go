package user

import (
	"miniblog/internal/pkg/core"
	"miniblog/internal/pkg/errno"
	"miniblog/internal/pkg/log"
	v1 "miniblog/pkg/api/miniblog/v1"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func (u *UserController) Create(c *gin.Context) {
	log.C(c).Infow("create user function called")

	var r v1.CreateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}

	if err := u.b.Users().Create(c, &r); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)

}
