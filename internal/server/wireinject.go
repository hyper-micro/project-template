package server

import (
	"github.com/google/wire"
	"github.com/hyper-micro/hyper/provider/http"
	"github.com/hyper-micro/hyper/provider/rpc"
	"github.com/hyper-micro/project-layout/internal/handler/restful"
	rpcHandler "github.com/hyper-micro/project-layout/internal/handler/rpc"
)

var ProviderSet = wire.NewSet(
	http.NewProvider,
	restful.NewGreeterRestfulHandler,
	wire.Struct(new(RestfulHandlerSet), "*"),

	rpc.NewProvider,
	rpcHandler.NewGreeterRpcServerHandler,
	wire.Struct(new(RpcHandlerSet), "*"),
)
