package store

import (
	"sync"

	"gorm.io/gorm"
)

var (
	S    *datastore
	once sync.Once
)

var _ IStore = (*datastore)(nil)

type IStore interface {
	Users() UserStore
}

type datastore struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *datastore {
	once.Do(func() {
		S = &datastore{
			db: db,
		}
	})
	return S
}

func (s *datastore) Users() UserStore {
	return newUsers(s.db)
}
