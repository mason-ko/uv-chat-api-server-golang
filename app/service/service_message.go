package service

import (
	"github.com/samber/lo"
	"uv-chat-api-server-golang/domain"
	"uv-chat-api-server-golang/domain/message"
	"uv-chat-api-server-golang/internal/appctx"
)

type messageService struct {
	repository domain.Repository
}

func (m *messageService) Create(ctx appctx.Context, msg message.Message) error {
	_, err := m.repository.MessageRepository().Create(msg.DBModel())
	return err
}

func (m *messageService) Delete(ctx appctx.Context, id uint) error {
	err := m.repository.MessageRepository().Delete(&message.DBMessageParam{
		ID: id,
	})
	return err
}

func (m *messageService) Get(ctx appctx.Context, id uint) (message.Message, error) {
	msg, err := m.repository.MessageRepository().Get(&message.DBMessageParam{
		ID: id,
	})
	if err != nil {
		return message.Message{}, err
	}
	return msg.Message(), nil
}

func (m *messageService) GetList(ctx appctx.Context, param message.GetListParam) ([]message.Message, error) {
	list, err := m.repository.MessageRepository().GetList(&message.DBMessageParam{}, param.Pagination, param.OrderBy)
	if err != nil {
		return nil, err
	}
	return lo.Map(list, func(item message.DBMessage, index int) message.Message {
		return item.Message()
	}), nil
}

func newMessageService(repository domain.Repository) message.Service {
	return &messageService{
		repository: repository,
	}
}
