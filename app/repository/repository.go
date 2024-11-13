package repository

import (
	"uv-chat-api-server-golang/domain"
	"uv-chat-api-server-golang/domain/message"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(newRepository),
)

type repository struct {
	external domain.External

	messageRepository domain.BaseRepository[message.Message]
}

func (r *repository) MessageRepository() domain.BaseRepository[message.Message] {
	return r.messageRepository
}

func newRepository(external domain.External) domain.Repository {
	return &repository{
		external:          external,
		messageRepository: newBaseRepository[message.Message](external),
	}
}
