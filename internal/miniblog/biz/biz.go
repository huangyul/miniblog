package biz

import (
	"miniblog/internal/miniblog/biz/user"
	"miniblog/internal/miniblog/store"
)

var _ IBiz = &biz{}

type IBiz interface {
	Users() user.UserBiz
}

type biz struct {
	ds store.IStore
}

func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}

// Users implements IBiz.
func (b *biz) Users() user.UserBiz {
	return user.New(b.ds)
}
