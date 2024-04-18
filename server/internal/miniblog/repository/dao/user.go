package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
)

var _ IUserDao = (*UserDao)(nil)

type IUserDao interface {
	Create(ctx context.Context, user User) error
}

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) IUserDao {
	return &UserDao{
		db: db,
	}
}

func (dao *UserDao) Create(ctx context.Context, user User) error {
	return dao.db.WithContext(ctx).Create(user).Error
}

type User struct {
	ID        int64     `gorm:"column:id;primary_key"`
	Password  string    `gorm:"column:password;not null"`
	Nickname  string    `gorm:"column:nickname"`
	Email     string    `gorm:"column:email;not null;unique"`
	Phone     string    `gorm:"column:phone;unique"`
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
}
