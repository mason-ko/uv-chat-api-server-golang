package service

import (
	"uv-chat-api-server-golang/domain"
	"uv-chat-api-server-golang/domain/message"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(newService),
)

type service struct {
	repository domain.Repository

	messageService message.Service
}

func (s *service) MessageService() message.Service {
	return s.messageService
}

func newService(repository domain.Repository) domain.Service {
	return &service{
		repository: repository,
	}
}
