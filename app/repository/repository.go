package repository

import (
	"uv-chat-api-server-golang/domain"
	"uv-chat-api-server-golang/domain/channel"
	"uv-chat-api-server-golang/domain/message"
	"uv-chat-api-server-golang/domain/user"

	"go.uber.org/fx"
)

var Modules = fx.Options(
	fx.Provide(newRepository),
)

type repository struct {
	external domain.External

	messageRepository domain.BaseRepository[message.DBMessage]
	userRepository    domain.BaseRepository[user.DBUser]
	channelRepository domain.BaseRepository[channel.DBChannel]
}

func (r *repository) UserRepository() domain.BaseRepository[user.DBUser] {
	return r.userRepository
}

func (r *repository) ChannelRepository() domain.BaseRepository[channel.DBChannel] {
	return r.channelRepository
}

func (r *repository) MessageRepository() domain.BaseRepository[message.DBMessage] {
	return r.messageRepository
}

func newRepository(external domain.External) domain.Repository {
	return &repository{
		external:          external,
		messageRepository: newBaseRepository[message.DBMessage](external),
		userRepository:    newBaseRepository[user.DBUser](external),
		channelRepository: newBaseRepository[channel.DBChannel](external),
	}
}
