package server

import (
	"github.com/google/wire"
	"github.com/hyper-micro/hyper/config"
	"github.com/hyper-micro/hyper/provider/rpc"
	rpcServer "github.com/hyper-micro/hyper/server/rpc"
	rpcHandler "github.com/hyper-micro/project-template/internal/handler/rpc"
	"github.com/hyper-micro/project-template/internal/repository"
	"github.com/hyper-micro/project-template/internal/service"
	"github.com/hyper-micro/project-template/internal/service/svcctx"
)

type RpcHandlerSet struct {
	Greeter *rpcHandler.GreeterRpcServerHandler
}

type RpcServer struct {
	rpcProvider rpc.Provider
}

func NewRpcProvider(conf config.Config) rpc.Provider {
	return rpc.NewProvider(
		conf,
		func(option *rpcServer.Option) {
			option.ServiceOpts = append(
				option.ServiceOpts,
				//grpc.ChainUnaryInterceptor(middleware.TracerUnaryServerInterceptor()),
			)
		},
	)
}

var RpcServerWireInject = wire.NewSet(
	repository.NewWaiterRepository,
	svcctx.NewGreeterServiceCtx,
	service.NewGreeterService,
	rpcHandler.NewGreeterRpcServerHandler,
	wire.Struct(new(RpcHandlerSet), "*"),
	NewRpcProvider,
)

func NewRpcServer(conf config.Config, srv rpc.Provider, handlerSet *RpcHandlerSet) *RpcServer {
	r := srv.Into()
	r.Handler(handlerSet.Greeter.RegisterService())
	return &RpcServer{rpcProvider: srv}
}

func (srv *RpcServer) Run() error {
	return srv.rpcProvider.Run()
}

func (srv *RpcServer) Shutdown() error {
	return srv.rpcProvider.Shutdown()
}

func (srv *RpcServer) Addr() string {
	return srv.rpcProvider.Addr()
}

func (srv *RpcServer) Name() string {
	return "RpcServer"
}
