package rpc

import (
	"context"

	"github.com/hyper-micro/hyper/server/rpc"
	greeter "github.com/hyper-micro/project-template/api/v1"
	"github.com/hyper-micro/project-template/internal/service"
	"google.golang.org/grpc"
)

type GreeterRpcServerHandler struct {
	greeter.UnimplementedGreeterServiceServer

	greeterSvc service.GreeterService
}

func NewGreeterRpcServerHandler(greeterSvc service.GreeterService) *GreeterRpcServerHandler {
	return &GreeterRpcServerHandler{
		greeterSvc: greeterSvc,
	}
}

func (srv *GreeterRpcServerHandler) RegisterService() rpc.HandlerFn {
	return func(s *grpc.Server) {
		greeter.RegisterGreeterServiceServer(s, srv)
	}
}

func (srv *GreeterRpcServerHandler) SayHello(ctx context.Context, req *greeter.SayHelloRequest) (*greeter.SayHelloReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	replyMessage, err := srv.greeterSvc.SayHello(ctx, req.GetWaiterId())
	if err != nil {
		return nil, err
	}

	reply := &greeter.SayHelloReply{Message: replyMessage}

	return reply, nil
}
