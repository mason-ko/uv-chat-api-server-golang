package domain

import (
	"uv-chat-api-server-golang/domain/channel"
	"uv-chat-api-server-golang/domain/message"
	"uv-chat-api-server-golang/domain/user"
)

type Controller interface {
	MessageController() message.Controller
	UserController() user.Controller
	ChannelController() channel.Controller
}
