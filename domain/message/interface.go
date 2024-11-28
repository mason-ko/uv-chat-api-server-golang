package message

import (
	"github.com/gin-gonic/gin"
	"uv-chat-api-server-golang/internal/appctx"
)

type Controller interface {
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetList(ctx *gin.Context)
}

type Service interface {
	Create(ctx appctx.Context, msg ReqCreateMessage) error
	Delete(ctx appctx.Context, id uint) error
	Get(ctx appctx.Context, id uint) (Message, error)
	GetList(ctx appctx.Context, param GetListParam) ([]Message, error)
}
