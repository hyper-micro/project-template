package svcctx

import (
	"context"

	"github.com/hyper-micro/project-template/internal/repository"
)

type GreeterServiceCtx struct {
	Ctx        context.Context
	WaiterRepo repository.WaiterRepository
}

func NewGreeterServiceCtx(ctx context.Context, waiterRepo repository.WaiterRepository) *GreeterServiceCtx {
	return &GreeterServiceCtx{
		Ctx:        ctx,
		WaiterRepo: waiterRepo,
	}
}
