package domain

import "uv-chat-api-server-golang/domain/message"

type Controller interface {
	MessageController() message.Controller
}
