package message

import (
	"uv-chat-api-server-golang/internal/ctx"
)

type Service interface {
	Create(ctx ctx.Context, msg Message) error
	Delete(ctx ctx.Context, id int) error
	Get(ctx ctx.Context, id int) (Message, error)
	GetList(ctx ctx.Context, param GetListParam) ([]Message, error)
}
