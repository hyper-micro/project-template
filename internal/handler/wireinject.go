package handler

import (
	"github.com/google/wire"
	"github.com/hyper-micro/project-layout/internal/service"
)

var ProviderSet = wire.NewSet(
	service.NewGreeterService,
)
