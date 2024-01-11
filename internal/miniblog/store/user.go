package store

import (
	"context"
	"gorm.io/gorm"
	"miniblog/internal/pkg/model"
)

// UserStore 定义了 user 模块在 store 层所实现的方法
type UserStore interface {
	// Create 新建 user 的方法
	Create(ctx context.Context, user *model.UserM) error
}

type Users struct {
	db *gorm.DB
}

var _ UserStore = &Users{}

func NewUser(db *gorm.DB) *Users {
	return &Users{db}
}

func (u *Users) Create(ctx context.Context, user *model.UserM) error {
	return u.db.Create(&user).Error
}
