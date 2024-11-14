package controller

import (
	"uv-chat-api-server-golang/domain"
	"uv-chat-api-server-golang/domain/message"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(newController),
)

type controller struct {
	service           domain.Service
	messageController message.Controller
}

func (c *controller) MessageController() message.Controller {
	return c.messageController
}

func newController(service domain.Service) domain.Controller {
	return &controller{
		service:           service,
		messageController: newMessageController(service),
	}
}
