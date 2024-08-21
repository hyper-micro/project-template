package restful

import (
	"fmt"

	"github.com/hyper-micro/hyper/server/web"
	"github.com/hyper-micro/project-template/internal/service"
)

type GreeterRestfulHandler struct {
	greeterSvc service.GreeterService
}

func NewGreeterRestfulHandler(greeterSvc service.GreeterService) *GreeterRestfulHandler {
	return &GreeterRestfulHandler{
		greeterSvc: greeterSvc,
	}
}

func (r *GreeterRestfulHandler) SayHello(ctx web.Ctx) {
	userId := ctx.QueryInt64("userId")
	replyMessage, err := r.greeterSvc.SayHello(ctx, userId)
	if err != nil {
		_ = ctx.String(
			fmt.Sprintf("Oops, err: %v", err),
		)
		return
	}

	_ = ctx.String(replyMessage)
}
