package svcctx

import (
	"github.com/google/wire"
	"github.com/hyper-micro/project-layout/internal/repository"
)

var ProviderSet = wire.NewSet(
	repository.NewAccountRepository,
)
