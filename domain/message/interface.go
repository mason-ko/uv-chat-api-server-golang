package message

import (
	"uv-chat-api-server-golang/internal/ctx"
)

type Service interface {
	Create(ctx ctx.Context, msg Message) error
	Delete(ctx ctx.Context, id uint) error
	Get(ctx ctx.Context, id uint) (Message, error)
	GetList(ctx ctx.Context, param GetListParam) ([]Message, error)
}
