package repository

import (
	"context"

	"github.com/huangyul/miniblog/internal/miniblog/domain"
	"github.com/huangyul/miniblog/internal/miniblog/repository/dao"
)

type IUserRepository interface {
	Create(ctx context.Context, user domain.User) error
}

type UserRepository struct {
	dao dao.IUserDao
}

func NewUserRepository(dao dao.IUserDao) IUserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) Create(ctx context.Context, user domain.User) error {
	return r.dao.Create(ctx, toEntiry(user))
}

func toEntiry(user domain.User) dao.User {
	return dao.User{
		Email:    user.Email,
		Password: user.Password,
	}
}
