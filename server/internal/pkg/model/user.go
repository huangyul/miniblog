package model

import "time"

type UserM struct {
	ID        int64  `gorm:"column:id;primaryKey"`
	Uesrname  string `gorm:"column:user_name;not null"`
	Password  string `gorm:"column:password;not null"`
	Nickname  string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *UserM) TableName() string {
	return "user"
}
