package repository

import (
	"context"

	"github.com/hyper-micro/project-template/internal/entity"
)

type Repository interface {
	Get(ctx context.Context, id int64) (*entity.Entity, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) Get(ctx context.Context, id int64) (*entity.Entity, error) {
	return &entity.Entity{
		ID:   id,
		Name: "Andy",
	}, nil
}
