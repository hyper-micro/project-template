package repository

import (
	"context"

	"github.com/hyper-micro/project-layout/internal/entity"
)

type AccountRepository interface {
	GetUser(ctx context.Context, userID int64) (*entity.User, error)
}

type accountRepository struct{}

func NewAccountRepository() AccountRepository {
	return &accountRepository{}
}

func (repo *accountRepository) GetUser(ctx context.Context, userID int64) (*entity.User, error) {
	userEntity := &entity.User{
		ID:   1,
		Name: "Nick",
	}
	return userEntity, nil
}
