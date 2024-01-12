package user

import (
	"context"
	"miniblog/internal/miniblog/store"
	"miniblog/internal/pkg/errno"
	"miniblog/internal/pkg/model"
	v1 "miniblog/pkg/api/miniblog/v1"
	"regexp"

	"github.com/jinzhu/copier"
)

type UserBiz interface {
	Create(ctx context.Context, r *v1.CreateUserRequest) error
}

type userBiz struct {
	ds store.IStore
}

func New(ds store.IStore) *userBiz {
	return &userBiz{
		ds: ds,
	}
}

// Create
func (b *userBiz) Create(ctx context.Context, r *v1.CreateUserRequest) error {
	var userM model.UserM
	_ = copier.Copy(&userM, r)

	if err := b.ds.Users().Create(ctx, &userM); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}
		return err
	}

	return nil
}

var _ UserBiz = &userBiz{}
