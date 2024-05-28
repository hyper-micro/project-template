package svcctx

import (
	"context"

	"github.com/hyper-micro/project-layout/internal/repository"
)

type GreeterServiceCtx struct {
	Ctx         context.Context
	AccountRepo repository.AccountRepository
}

func NewGreeterServiceCtx(ctx context.Context, accountRepo repository.AccountRepository) *GreeterServiceCtx {
	return &GreeterServiceCtx{
		Ctx:         ctx,
		AccountRepo: accountRepo,
	}
}
