//go:build wireinject
// +build wireinject

package wireinject

import (
	"context"

	"github.com/google/wire"
	"github.com/hyper-micro/hyper/config"
	"github.com/hyper-micro/project-template/internal/handler/restful"
	rpcHandler "github.com/hyper-micro/project-template/internal/handler/rpc"
	"github.com/hyper-micro/project-template/internal/repository"
	"github.com/hyper-micro/project-template/internal/server"
	"github.com/hyper-micro/project-template/internal/service"
	"github.com/hyper-micro/project-template/internal/service/svcctx"
)

var serviceSets = wire.NewSet(
	repository.NewRepository,
	svcctx.NewServiceCtx,
	service.NewService,
)

var httpServerSets = wire.NewSet(
	restful.NewHandler,
	wire.Struct(new(server.RestfulHandlerSet), "*"),
	server.NewHttpProvider,
	server.NewHttpServer,
)

var rpcServerSets = wire.NewSet(
	rpcHandler.NewHandler,
	wire.Struct(new(server.RpcHandlerSet), "*"),
	server.NewRpcProvider,
	server.NewRpcServer,
)

func NewHttpServer(ctx context.Context, conf config.Config) (*server.HttpServer, func(), error) {
	wire.Build(
		serviceSets,
		httpServerSets,
	)
	return &server.HttpServer{}, nil, nil
}

func NewRpcServer(ctx context.Context, conf config.Config) (*server.RpcServer, func(), error) {
	wire.Build(
		serviceSets,
		rpcServerSets,
	)
	return &server.RpcServer{}, nil, nil
}

type Servers struct {
	HttpServer *server.HttpServer
	RpcServer  *server.RpcServer
}

func NewServer(ctx context.Context, conf config.Config) (*Servers, func(), error) {
	wire.Build(
		serviceSets,
		httpServerSets,
		rpcServerSets,
		wire.Struct(new(Servers), "*"),
	)

	return &Servers{}, nil, nil
}
