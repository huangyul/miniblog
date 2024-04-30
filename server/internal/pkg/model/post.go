package model

import "time"

type PostM struct {
	ID        int64  `gorm:"column:id;primaryKey"`
	Username  string `gorm:"not null"`
	PostID    int64  `gorm:"column:post_id;"`
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *PostM) TableName() string {
	return "post"
}
