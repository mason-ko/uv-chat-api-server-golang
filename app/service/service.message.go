package service

import (
	"uv-chat-api-server-golang/domain"
	"uv-chat-api-server-golang/domain/message"
	"uv-chat-api-server-golang/internal/ctx"
)

type messageService struct {
	repository domain.Repository
}

func (m *messageService) Create(ctx ctx.Context, msg message.Message) error {
	//TODO implement me
	panic("implement me")
}

func (m *messageService) Delete(ctx ctx.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (m *messageService) Get(ctx ctx.Context, id int) (message.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (m *messageService) GetList(ctx ctx.Context, param message.GetListParam) ([]message.Message, error) {
	//TODO implement me
	panic("implement me")
}

func newMessageService(repository domain.Repository) message.Service {
	return &messageService{
		repository: repository,
	}
}
