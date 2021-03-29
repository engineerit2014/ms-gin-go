package services

import (
	"context"
	"github.com/laironacosta/ms-gin-go/controllers/dto"
	repo "github.com/laironacosta/ms-gin-go/repository"
)

type UserServiceInterface interface {
	Create(ctx context.Context, user dto.User) error
	GetByEmail(ctx context.Context, email string) (*dto.User, error)
	UpdateByEmail(ctx context.Context, email string) error
	DeleteByEmail(ctx context.Context, email string) error
}

type UserService struct {
	userRepo repo.UserRepositoryInterface
}

func NewUserService(userRepo repo.UserRepositoryInterface) UserServiceInterface {
	return &UserService{
		userRepo,
	}
}

func (s *UserService) Create(ctx context.Context, user dto.User) error {
	if err := s.userRepo.Create(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*dto.User, error) {
	u, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return &dto.User{}, err
	}
	return u, nil
}

func (s *UserService) UpdateByEmail(ctx context.Context, email string) error {
	if err := s.userRepo.UpdateByEmail(ctx, email); err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteByEmail(ctx context.Context, email string) error {
	if err := s.userRepo.UpdateByEmail(ctx, email); err != nil {
		return err
	}
	return nil
}
