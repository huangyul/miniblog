package store

import (
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
	S    *database
)

var _ IStore = &database{}

type IStore interface {
	User() UserStore
}

type database struct {
	db *gorm.DB
}

func (d *database) User() UserStore {
	return newUsers(d.db)
}

func NewStore(db *gorm.DB) *database {
	once.Do(func() {
		S = &database{db: db}
	})

	return S
}
