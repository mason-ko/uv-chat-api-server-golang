package domain

import (
	"gorm.io/gorm"
	"uv-chat-api-server-golang/domain/message"
)

type Repository interface {
	MessageRepository() BaseRepository[message.Message]
}

type BaseRepository[T ModelWithID] interface {
	Create(t T) (uint, error)
	Delete(BaseWhereModel) error
	Update(model BaseWhereModel, t T) error
	Get(BaseWhereModel) (T, error)
	GetList(BaseWhereModel) ([]T, error)
}

type ModelWithID interface {
	GetID() uint
}

type BaseWhereModel interface {
	SetExpression() func(tx *gorm.DB) *gorm.DB
}
