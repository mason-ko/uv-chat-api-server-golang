package repository

import (
	"uv-chat-api-server-golang/domain"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(newRepository),
)

type repository struct {
	external domain.External
}

func newRepository(external domain.External) domain.Repository {
	return &repository{
		external: external,
	}
}
