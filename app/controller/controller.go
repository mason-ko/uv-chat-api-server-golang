package controller

import (
	"uv-chat-api-server-golang/domain"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(newController),
)

type controller struct {
	servcie domain.Service
}

func newController(servcie domain.Service) domain.Controller {
	return &controller{
		servcie: servcie,
	}
}
