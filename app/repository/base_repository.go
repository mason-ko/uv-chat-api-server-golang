package repository

import (
	"uv-chat-api-server-golang/domain"
)

type baseRepository[T domain.ModelWithID] struct {
	external domain.External
}

func (b *baseRepository[T]) Create(t T) (uint, error) {
	result := b.external.DB().Create(&t)
	return t.GetID(), result.Error
}

func (b *baseRepository[T]) Delete(model domain.BaseWhereModel) error {
	return model.SetExpression()(b.external.DB()).Delete(new(T)).Error
}

func (b *baseRepository[T]) Update(model domain.BaseWhereModel, t T) error {
	return model.SetExpression()(b.external.DB()).Updates(t).Error
}

func (b *baseRepository[T]) Get(model domain.BaseWhereModel) (T, error) {
	var ret T
	err := model.SetExpression()(b.external.DB()).First(&ret).Error
	return ret, err
}

func (b *baseRepository[T]) GetList(model domain.BaseWhereModel) ([]T, error) {
	var ret []T
	err := model.SetExpression()(b.external.DB()).Find(&ret).Error
	return ret, err
}

func newBaseRepository[T domain.ModelWithID](external domain.External) domain.BaseRepository[T] {
	return &baseRepository[T]{
		external: external,
	}
}
