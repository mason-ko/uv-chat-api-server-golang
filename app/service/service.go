package service

import (
	"uv-chat-api-server-golang/domain"
	"uv-chat-api-server-golang/domain/channel"
	"uv-chat-api-server-golang/domain/message"
	"uv-chat-api-server-golang/domain/user"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(newService),
)

type service struct {
	repository domain.Repository

	messageService message.Service
	userService    user.Service
	channelService channel.Service
}

func (s *service) ChannelService() channel.Service {
	return s.channelService
}

func (s *service) UserService() user.Service {
	return s.userService
}

func (s *service) MessageService() message.Service {
	return s.messageService
}

func newService(repository domain.Repository) domain.Service {
	return &service{
		repository:     repository,
		messageService: newMessageService(repository),
		userService:    newUserService(repository),
		channelService: newChannelService(repository),
	}
}
