package service

import (
	"github.com/go-redis/redis/v8"
	"github.com/samber/lo"
	"strconv"
	"time"
	"uv-chat-api-server-golang/domain"
	"uv-chat-api-server-golang/domain/channel"
	"uv-chat-api-server-golang/domain/message"
	"uv-chat-api-server-golang/domain/user"
	"uv-chat-api-server-golang/internal/appctx"
	"uv-chat-api-server-golang/internal/util"
)

type messageService struct {
	repository domain.Repository

	redisClient *redis.Client
}

func (m *messageService) Create(ctx appctx.Context, req message.ReqCreateMessage) error {
	// 해당 채널의 다른 user 정보 찾음
	chUsers, err := m.repository.ChannelUsersRepository().GetList(&channel.ChannelUsersParam{
		ChannelID: req.ChannelID,
	}, nil, nil)
	if err != nil {
		return err
	}
	var otherUserID uint
	for _, user := range chUsers {
		if user.UserID != req.UserID {
			otherUserID = user.UserID
		}
	}
	// 해당 국가 정보로 api 발송
	otherUser, err := m.repository.UserRepository().Get(&user.DBUserParam{ID: otherUserID})
	if err != nil {
		return err
	}

	//SendTranslateHttp
	translatedContent := util.SendTranslateHttp(req.Content, otherUser.Country)
	dbMsg := message.DBMessage{
		ChannelID:         req.ChannelID,
		UserID:            req.UserID,
		Content:           req.Content,
		TranslatedContent: translatedContent,
	}
	id, err := m.repository.MessageRepository().Create(dbMsg)
	if err != nil {
		return err
	}
	dbMsg.ID = id
	dbMsg.CreatedAt = time.Now()

	// send socket
	util.SendSocketMessage(m.redisClient, []string{"channel:" + strconv.Itoa(int(req.ChannelID))}, "create_message", dbMsg.Message())

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

func newMessageService(repository domain.Repository, external domain.External) message.Service {
	return &messageService{
		repository:  repository,
		redisClient: external.RedisClient(),
	}
}
