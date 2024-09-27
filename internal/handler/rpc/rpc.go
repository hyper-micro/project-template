package rpc

import (
	"context"

	"github.com/hyper-micro/hyper/server/rpc"
	apiService "github.com/hyper-micro/project-template/api/v1"
	"github.com/hyper-micro/project-template/internal/service"
	"google.golang.org/grpc"
)

type Handler struct {
	apiService.UnimplementedProjectTemplateServiceServiceServer
	svc service.Service
}

func NewHandler(svc service.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (srv *Handler) RegisterService() rpc.HandlerFn {
	return func(s *grpc.Server) {
		apiService.RegisterProjectTemplateServiceServiceServer(s, srv)
	}
}

func (srv *Handler) SayHello(ctx context.Context, req *apiService.ProjectTemplateServiceRequest) (*apiService.ProjectTemplateServiceReply, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	replyMessage, err := srv.svc.Hello(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	reply := &apiService.ProjectTemplateServiceReply{Message: replyMessage}

	return reply, nil
}
