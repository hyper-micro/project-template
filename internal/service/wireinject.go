package service

import (
	"github.com/google/wire"
	"github.com/hyper-micro/project-layout/internal/service/svcctx"
)

var ProviderSet = wire.NewSet(
	svcctx.NewGreeterServiceCtx,
)
