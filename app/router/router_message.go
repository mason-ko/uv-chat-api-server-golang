package router

import (
	"github.com/gin-gonic/gin"
	"uv-chat-api-server-golang/domain"
)

func registerMessageRouter(router *gin.Engine, controller domain.Controller) {
	g := router.Group("/api/messages")
	g.GET("/:id", controller.MessageController().Get)
	g.GET("/", controller.MessageController().GetList)
	g.DELETE("/:id", controller.MessageController().Delete)
	g.POST("/", controller.MessageController().Create)
}
