package svcctx

import (
	"context"

	"github.com/hyper-micro/project-template/internal/repository"
)

type ServiceCtx struct {
	Ctx  context.Context
	Repo repository.Repository
}

func NewServiceCtx(ctx context.Context, repo repository.Repository) *ServiceCtx {
	return &ServiceCtx{
		Ctx:  ctx,
		Repo: repo,
	}
}
