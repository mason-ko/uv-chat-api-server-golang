package domain

import (
	"gorm.io/gorm"
	"uv-chat-api-server-golang/domain/message"
	"uv-chat-api-server-golang/internal/common"
)

type Repository interface {
	MessageRepository() BaseRepository[message.Message]
}

type BaseRepository[T ModelWithID] interface {
	Create(t T) (uint, error)
	Delete(model BaseWhereModel) error
	Update(model BaseWhereModel, t T) error
	Get(model BaseWhereModel) (T, error)
	GetList(model BaseWhereModel, pagination *common.Pagination, orderBy *common.OrderBy) ([]T, error)
}

type ModelWithID interface {
	GetID() uint
}

type BaseWhereModel interface {
	SetExpression() func(tx *gorm.DB) *gorm.DB
}
