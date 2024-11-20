package router

import (
	"github.com/gin-gonic/gin"
	"uv-chat-api-server-golang/domain"
)

func registerChannelRouter(router *gin.Engine, controller domain.Controller) {
	g := router.Group("/api/channels")
	g.GET("/:id", controller.ChannelController().Get)
	g.GET("/", controller.ChannelController().GetList)
	g.DELETE("/:id", controller.ChannelController().Delete)
	g.POST("/", controller.ChannelController().Create)
}
