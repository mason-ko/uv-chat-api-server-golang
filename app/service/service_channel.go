package service

import (
	"github.com/samber/lo"
	"uv-chat-api-server-golang/domain"
	"uv-chat-api-server-golang/domain/channel"
	"uv-chat-api-server-golang/domain/message"
	"uv-chat-api-server-golang/internal/appctx"
)

type channelService struct {
	repository domain.Repository
}

func (m *channelService) Create(ctx appctx.Context, msg channel.Channel) error {
	_, err := m.repository.ChannelRepository().Create(msg.DBModel())
	return err
}

func (m *channelService) Delete(ctx appctx.Context, id uint) error {
	err := m.repository.ChannelRepository().Delete(&message.DBMessageParam{
		ID: id,
	})
	return err
}

func (m *channelService) Get(ctx appctx.Context, id uint) (channel.Channel, error) {
	msg, err := m.repository.ChannelRepository().Get(&message.DBMessageParam{
		ID: id,
	})
	if err != nil {
		return channel.Channel{}, err
	}
	return msg.Channel(), nil
}

func (m *channelService) GetList(ctx appctx.Context, param channel.GetListParam) ([]channel.Channel, error) {
	list, err := m.repository.ChannelRepository().GetList(&message.DBMessageParam{}, param.Pagination, param.OrderBy)
	if err != nil {
		return nil, err
	}
	return lo.Map(list, func(item channel.DBChannel, index int) channel.Channel {
		return item.Channel()
	}), nil
}

func newChannelService(repository domain.Repository) channel.Service {
	return &channelService{
		repository: repository,
	}
}
