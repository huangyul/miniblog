package store

import (
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
	s    *datastore
)

type IStore interface {
	Users() UserStore
}

type datastore struct {
	db *gorm.DB
}

func (d *datastore) Users() UserStore {
	return NewUser(d.db)
}

func NewStore(db *gorm.DB) *datastore {
	once.Do(func() {
		s = &datastore{
			db,
		}
	})
	return s
}

var _ IStore = &datastore{}
