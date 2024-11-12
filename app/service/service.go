package service

import (
	"uv-chat-api-server-golang/domain"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(newService),
)

type service struct {
	repository domain.Repository
}

func newService(repository domain.Repository) domain.Service {
	return &service{
		repository: repository,
	}
}
