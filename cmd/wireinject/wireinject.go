//go:build wireinject
// +build wireinject

package wireinject

import (
	"context"

	"github.com/google/wire"
	"github.com/hyper-micro/hyper/config"
	"github.com/hyper-micro/project-layout/internal/handler"
	"github.com/hyper-micro/project-layout/internal/server"
	"github.com/hyper-micro/project-layout/internal/service"
	"github.com/hyper-micro/project-layout/internal/service/svcctx"
)

func NewHttpServer(ctx context.Context, conf config.Config) (*server.HttpServer, func(), error) {
	wire.Build(
		svcctx.ProviderSet,
		service.ProviderSet,
		handler.ProviderSet,
		server.ProviderSet,
		server.NewHttpServer,
	)
	return &server.HttpServer{}, nil, nil
}

func NewRpcServer(ctx context.Context, conf config.Config) (*server.RpcServer, func(), error) {
	wire.Build(
		svcctx.ProviderSet,
		service.ProviderSet,
		handler.ProviderSet,
		server.ProviderSet,
		server.NewRpcServer,
	)
	return &server.RpcServer{}, nil, nil
}
