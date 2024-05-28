package service

import (
	"context"
	"fmt"

	"github.com/hyper-micro/project-layout/internal/service/svcctx"
)

type GreeterService interface {
	SayHello(ctx context.Context, userID int64) (string, error)
}

type greeterService struct {
	svcCtx *svcctx.GreeterServiceCtx
}

func NewGreeterService(svcCtx *svcctx.GreeterServiceCtx) GreeterService {
	return &greeterService{
		svcCtx: svcCtx,
	}
}

func (s *greeterService) SayHello(ctx context.Context, userID int64) (string, error) {
	userEntity, err := s.svcCtx.AccountRepo.GetUser(ctx, userID)
	if err != nil {
		return "", err
	}

	replyMessage := fmt.Sprintf("Hello, %s", userEntity.Name)

	return replyMessage, nil
}
