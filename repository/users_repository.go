package repository

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/laironacosta/ms-gin-go/controllers/dto"
)

type UserRepositoryInterface interface {
	Create(ctx context.Context, user dto.User) error
	GetByEmail(ctx context.Context, email string) (*dto.User, error)
	UpdateByEmail(ctx context.Context, email string) error
	DeleteByEmail(ctx context.Context, email string) error
}

type UserRepository struct {
	DB *pg.DB
}

func NewUserRepository(db *pg.DB) UserRepositoryInterface {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user dto.User) error {
	return nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*dto.User, error) {
	return &dto.User{}, nil
}

func (r *UserRepository) UpdateByEmail(ctx context.Context, email string) error {
	return nil
}

func (r *UserRepository) DeleteByEmail(ctx context.Context, email string) error {
	return nil
}
