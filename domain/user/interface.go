package user

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
	Create(ctx appctx.Context, u User) error
	Delete(ctx appctx.Context, id uint) error
	Get(ctx appctx.Context, id uint) (User, error)
	GetList(ctx appctx.Context, param GetListParam) ([]User, error)
}
