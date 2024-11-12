package repository

import (
	"uv-chat-api-server-golang/domain"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(newRepository),
)

type repository struct {
}

func newRepository() domain.Repository {
	return &repository{}
}
