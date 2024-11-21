package router

import (
	"github.com/gin-gonic/gin"
	"uv-chat-api-server-golang/domain"
)

func registerUserRouter(router *gin.Engine, controller domain.Controller) {
	g := router.Group("/api/users")
	g.POST("", controller.UserController().Create)
	g.GET("", controller.UserController().GetList)
}
