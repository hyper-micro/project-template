package server

import (
	"github.com/google/wire"
	"github.com/hyper-micro/hyper/config"
	"github.com/hyper-micro/hyper/provider/http"
	"github.com/hyper-micro/hyper/server/web"
	"github.com/hyper-micro/project-template/internal/handler/restful"
	"github.com/hyper-micro/project-template/internal/repository"
	"github.com/hyper-micro/project-template/internal/service"
	"github.com/hyper-micro/project-template/internal/service/svcctx"
)

type RestfulHandlerSet struct {
	Greeter *restful.GreeterRestfulHandler
}

type HttpServer struct {
	httpProvider http.Provider
}

func NewHttpProvider(conf config.Config) http.Provider {
	return http.NewProvider(
		conf,
		func(option *web.Option) {},
	)
}

var HttpServerWireInject = wire.NewSet(
	repository.NewWaiterRepository,
	svcctx.NewGreeterServiceCtx,
	service.NewGreeterService,
	restful.NewGreeterRestfulHandler,
	wire.Struct(new(RestfulHandlerSet), "*"),
	NewHttpProvider,
)

func NewHttpServer(cfg config.Config, srv http.Provider, handlerSet *RestfulHandlerSet) *HttpServer {
	r := srv.Into()
	r.Get("/sayHello", handlerSet.Greeter.SayHello)

	return &HttpServer{httpProvider: srv}
}

func (srv *HttpServer) Run() error {
	return srv.httpProvider.Run()
}

func (srv *HttpServer) Shutdown() error {
	return srv.httpProvider.Shutdown()
}

func (srv *HttpServer) Addr() string {
	return srv.httpProvider.Addr()
}

func (srv *HttpServer) Name() string {
	return "HttpServer"
}
