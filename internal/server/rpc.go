package server

import (
	"github.com/hyper-micro/hyper/config"
	"github.com/hyper-micro/hyper/provider/rpc"
	rpcHandler "github.com/hyper-micro/project-layout/internal/handler/rpc"
)

type RpcHandlerSet struct {
	Greeter *rpcHandler.GreeterRpcServerHandler
}

type RpcServer struct {
	rpcProvider rpc.Provider
}

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
