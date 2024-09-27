package service

import (
	"context"
	"fmt"

	"github.com/hyper-micro/project-template/internal/service/svcctx"
)

type Service interface {
	Hello(ctx context.Context, id int64) (string, error)
}

type service struct {
	svcCtx *svcctx.ServiceCtx
}

func NewService(svcCtx *svcctx.ServiceCtx) Service {
	return &service{
		svcCtx: svcCtx,
	}
}

func (s *service) Hello(ctx context.Context, id int64) (string, error) {
	entity, err := s.svcCtx.Repo.Get(ctx, id)
	if err != nil {
		return "", err
	}

	replyMessage := fmt.Sprintf("Hello, I'm %s, happy to serve you.", entity.Name)

	return replyMessage, nil
}
