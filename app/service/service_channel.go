package service

import (
	"errors"
	"github.com/samber/lo"
	"uv-chat-api-server-golang/domain"
	"uv-chat-api-server-golang/domain/channel"
	"uv-chat-api-server-golang/domain/message"
	"uv-chat-api-server-golang/internal/appctx"
)

type channelService struct {
	repository domain.Repository
}

func (m *channelService) Create(ctx appctx.Context, req channel.ReqCreateChannel) (uint, error) {
	if req.UserID == req.TargetUserID {
		return 0, errors.New("not yet")
	}

	channelUsers, err := m.repository.ChannelUsersRepository().GetList(&channel.ChannelUsersParam{
		UserIDs: []int{req.UserID, req.TargetUserID},
	}, nil, nil)
	if err != nil {
		return 0, err
	}

	if len(channelUsers) > 0 {
		return channelUsers[0].ChannelID, nil
	}

	// 채널 생성
	// 채널 Users 생성
	chId, err := m.repository.ChannelRepository().Create(channel.DBChannel{})
	if err != nil {
		return 0, err
	}
	// todo 배치
	for _, userID := range []int{req.UserID, req.TargetUserID} {
		_, err = m.repository.ChannelUsersRepository().Create(channel.DBChannelUsers{
			ChannelID: chId,
			UserID:    uint(userID),
		})
		if err != nil {
			return 0, err
		}
	}
	return chId, err
}

func (m *channelService) Delete(ctx appctx.Context, id uint) error {
	err := m.repository.ChannelRepository().Delete(&channel.ChannelParam{
		ID: id,
	})
	if err != nil {
		return err
	}
	err = m.repository.ChannelUsersRepository().Delete(&channel.ChannelUsersParam{
		ChannelID: id,
	})
	if err != nil {
		return err
	}
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
