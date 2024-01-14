package user

import (
	"context"
	"github.com/jinzhu/copier"
	"miniblog/internal/miniblog/store"
	"miniblog/internal/pkg/errno"
	"miniblog/internal/pkg/model"
	v1 "miniblog/pkg/api/miniblog/v1"
	"regexp"
)

var _ UserBiz = &userBiz{}

type UserBiz interface {
	Create(ctx context.Context, req *v1.CreateUserRequest) error
}

type userBiz struct {
	ds store.IStore
}

func (u *userBiz) Create(ctx context.Context, req *v1.CreateUserRequest) error {
	var userM model.UserM
	_ = copier.Copy(&userM, req)

	if err := u.ds.User().Create(ctx, &userM); err != nil {
		if match, _ := regexp.MatchString("Dulicate entry '.*' for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}

		return err
	}

	return nil
}

func New(ds store.IStore) *userBiz {
	return &userBiz{ds}
}
