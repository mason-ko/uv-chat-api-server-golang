package domain

import (
	"uv-chat-api-server-golang/domain/channel"
	"uv-chat-api-server-golang/domain/message"
	"uv-chat-api-server-golang/domain/user"
)

type Service interface {
	ChannelService() channel.Service
	MessageService() message.Service
	UserService() user.Service
}
