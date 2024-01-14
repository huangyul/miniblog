package store

import (
	"context"
	"gorm.io/gorm"
	"miniblog/internal/pkg/model"
)

var _ UserStore = &users{}

type UserStore interface {
	Create(ctx context.Context, user *model.UserM) error
}

type users struct {
	db *gorm.DB
}

func (u *users) Create(ctx context.Context, user *model.UserM) error {
	return u.db.Create(&user).Error
}

func newUsers(db *gorm.DB) *users {
	return &users{
		db: db,
	}
}
