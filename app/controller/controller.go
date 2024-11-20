package controller

import (
	"uv-chat-api-server-golang/domain"
	"uv-chat-api-server-golang/domain/channel"
	"uv-chat-api-server-golang/domain/message"
	"uv-chat-api-server-golang/domain/user"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(newController),
)

type controller struct {
	service           domain.Service
	messageController message.Controller
	userController    user.Controller
	channelController channel.Controller
}

func (c *controller) UserController() user.Controller {
	return c.userController
}

func (c *controller) ChannelController() channel.Controller {
	return c.channelController
}

func (c *controller) MessageController() message.Controller {
	return c.messageController
}

func newController(service domain.Service) domain.Controller {
	return &controller{
		service:           service,
		messageController: newMessageController(service),
		userController:    newUserController(service),
		channelController: newChannelController(service),
	}
}
