package server

import (
	"github.com/hyper-micro/hyper/config"
	"github.com/hyper-micro/hyper/provider/http"
	"github.com/hyper-micro/project-layout/internal/handler/restful"
)

type RestfulHandlerSet struct {
	Greeter *restful.GreeterRestfulHandler
}

type HttpServer struct {
	httpProvider http.Provider
}

func NewHttpServer(cfg config.Config, srv http.Provider, handlerSet *RestfulHandlerSet) *HttpServer {
	r := srv.Into()
	r.Get("/say-hello", handlerSet.Greeter.SayHello)

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
