//go:build wireinject
// +build wireinject

package wireinject

import (
	"context"

	"github.com/google/wire"
	"github.com/hyper-micro/hyper/config"
	"github.com/hyper-micro/project-template/internal/server"
)

func NewHttpServer(ctx context.Context, conf config.Config) (*server.HttpServer, func(), error) {
	wire.Build(
		server.HttpServerWireInject,
		server.NewHttpServer,
	)
	return &server.HttpServer{}, nil, nil
}

func NewRpcServer(ctx context.Context, conf config.Config) (*server.RpcServer, func(), error) {
	wire.Build(
		server.RpcServerWireInject,
		server.NewRpcServer,
	)
	return &server.RpcServer{}, nil, nil
}
