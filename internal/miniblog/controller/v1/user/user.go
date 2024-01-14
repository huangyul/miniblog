package user

import (
	"miniblog/internal/miniblog/biz"
	"miniblog/internal/miniblog/store"
)

type UserController struct {
	b biz.IBiz
}

func New(ds store.IStore) *UserController {
	return &UserController{b: biz.NewBiz(ds)}
}
