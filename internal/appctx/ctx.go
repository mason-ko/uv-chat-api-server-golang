package appctx

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Context struct {
	context.Context

	UserID int
}

func NewContext(ctx *gin.Context) Context {
	return Context{
		Context: ctx.Request.Context(),
		//UserIDs: 추후
	}
}
