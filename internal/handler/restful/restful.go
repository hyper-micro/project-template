package restful

import (
	"fmt"

	"github.com/hyper-micro/hyper/server/web"
	"github.com/hyper-micro/project-template/internal/service"
)

type Handler struct {
	svc service.Service
}

func NewHandler(svc service.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (r *Handler) SayHello(ctx web.Ctx) {
	id := ctx.QueryInt64("id")
	replyMessage, err := r.svc.Hello(ctx, id)
	if err != nil {
		_ = ctx.String(
			fmt.Sprintf("Oops, err: %v", err),
		)
		return
	}

	_ = ctx.String(replyMessage)
}
