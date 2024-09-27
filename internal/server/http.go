package server

import (
	"github.com/hyper-micro/hyper/config"
	"github.com/hyper-micro/hyper/provider/http"
	"github.com/hyper-micro/hyper/server/web"
	"github.com/hyper-micro/project-template/internal/handler/restful"
)

type RestfulHandlerSet struct {
	Handler *restful.Handler
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

func NewHttpServer(cfg config.Config, srv http.Provider, handlerSet *RestfulHandlerSet) *HttpServer {
	r := srv.Into()
	r.Get("/hello", handlerSet.Handler.SayHello)

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
