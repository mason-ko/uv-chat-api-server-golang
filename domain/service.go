package domain

import "uv-chat-api-server-golang/domain/message"

type Service interface {
	//ChannelService() channel.Service
	MessageService() message.Service
}
