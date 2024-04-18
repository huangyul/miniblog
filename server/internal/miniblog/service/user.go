package service

import (
	"context"

	"github.com/huangyul/miniblog/internal/miniblog/domain"
	"github.com/huangyul/miniblog/internal/miniblog/repository"
)

var _ IUserService = (*UserService)(nil)

type IUserService interface {
	Create(ctx context.Context, email string, password string) error
}

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) IUserService {
	return &UserService{
		repo: repo,
	}
}

// Create implements IUserService.
func (u *UserService) Create(ctx context.Context, email string, password string) error {
	return u.repo.Create(ctx, domain.User{
		Email:    email,
		Password: password,
	})
}
