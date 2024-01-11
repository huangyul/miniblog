package store

import (
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
	s    *Datastore
)

type IStore interface {
	Users() UserStore
}

type Datastore struct {
	db *gorm.DB
}

func (d *Datastore) Users() UserStore {
	return NewUser(d.db)
}

func NewStore(db *gorm.DB) *Datastore {
	once.Do(func() {
		s = &Datastore{
			db,
		}
	})
	return s
}

var _ IStore = &Datastore{}
