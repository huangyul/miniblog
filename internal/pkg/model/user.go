package model

import "time"

type UserM struct {
	ID        int64     `gorm:"column:id;primary_key"`
	Username  string    `gorm:"column:username;not null"`
	Password  string    `gorm:"column:password;not null"`
	Nickname  string    `gorm:"column:nickname"`
	Email     string    `gorm:"column:email"`
	Phone     string    `gorm:"column:phone"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (u *UserM) TableName() string {
	return "user"
}
