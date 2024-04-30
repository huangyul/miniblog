package store

import (
	"context"

	"github.com/huangyul/miniblog/internal/pkg/model"
	"gorm.io/gorm"
)

var _ UserStore = (*users)(nil)

type UserStore interface {
	Create(ctx context.Context, user *model.UserM) error
}

type users struct {
	db *gorm.DB
}

func newUsers(db *gorm.DB) *users {
	return &users{db: db}
}

func (u *users) Create(ctx context.Context, user *model.UserM) error {
	return u.db.WithContext(ctx).Create(&user).Error
}
