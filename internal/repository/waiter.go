package repository

import (
	"context"
	"errors"

	"github.com/hyper-micro/project-template/internal/entity"
)

type WaiterRepository interface {
	Get(ctx context.Context, userID int64) (*entity.Waiter, error)
}

type waiterRepository struct{}

func NewWaiterRepository() WaiterRepository {
	return &waiterRepository{}
}

var users = map[int64]*entity.Waiter{
	1: {
		ID:   1,
		Name: "Tom",
	},
	2: {
		ID:   2,
		Name: "Tony",
	},
	3: {
		ID:   3,
		Name: "Nick",
	},
}

func (repo *waiterRepository) Get(ctx context.Context, userID int64) (*entity.Waiter, error) {
	userEntity, ok := users[userID]
	if !ok {
		return nil, errors.New("user not exist")
	}
	return userEntity, nil
}
