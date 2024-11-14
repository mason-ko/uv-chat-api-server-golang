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

	messageRepository domain.BaseRepository[message.DBMessage]
}

func (r *repository) MessageRepository() domain.BaseRepository[message.DBMessage] {
	return r.messageRepository
}

func newRepository(external domain.External) domain.Repository {
	return &repository{
		external:          external,
		messageRepository: newBaseRepository[message.DBMessage](external),
	}
}
