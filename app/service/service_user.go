package service

import (
	"github.com/samber/lo"
	"uv-chat-api-server-golang/domain"
	"uv-chat-api-server-golang/domain/message"
	"uv-chat-api-server-golang/domain/user"
	"uv-chat-api-server-golang/internal/appctx"
)

type userService struct {
	repository domain.Repository
}

func (m *userService) Create(ctx appctx.Context, msg user.User) error {
	if msg.ID != 0 {
		_, err := m.Get(ctx, msg.ID)
		if err == nil {
			//update
			return m.repository.UserRepository().Update(&user.DBUserParam{
				ID: msg.ID,
			}, msg.DBModel())
		}
	}

	_, err := m.repository.UserRepository().Create(msg.DBModel())
	return err
}

func (m *userService) Delete(ctx appctx.Context, id uint) error {
	err := m.repository.UserRepository().Delete(&message.DBMessageParam{
		ID: id,
	})
	return err
}

func (m *userService) Get(ctx appctx.Context, id uint) (user.User, error) {
	msg, err := m.repository.UserRepository().Get(&message.DBMessageParam{
		ID: id,
	})
	if err != nil {
		return user.User{}, err
	}
	return msg.User(), nil
}

func (m *userService) GetList(ctx appctx.Context, param user.GetListParam) ([]user.User, error) {
	list, err := m.repository.UserRepository().GetList(&message.DBMessageParam{}, param.Pagination, param.OrderBy)
	if err != nil {
		return nil, err
	}
	return lo.Map(list, func(item user.DBUser, index int) user.User {
		return item.User()
	}), nil
}

func newUserService(repository domain.Repository) user.Service {
	return &userService{
		repository: repository,
	}
}
