package controller

import (
	"uv-chat-api-server-golang/domain"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(newController),
)

type controller struct {
	service domain.Service
}

func newController(service domain.Service) domain.Controller {
	return &controller{
		service: service,
	}
}
