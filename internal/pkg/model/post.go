package model

import "time"

type PostM struct {
	ID        int64     `gorm:"column:id;primary_key"`
	Username  string    `gorm:"column:username;not null"`
	PostID    string    `gorm:"column:post_id;not null"`
	Title     string    `gorm:"column:title;not null"`
	Content   string    `gorm:"column:content"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (p *PostM) TableName() string {
	return "post"
}
