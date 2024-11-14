package repository

import (
	"uv-chat-api-server-golang/domain"
	"uv-chat-api-server-golang/internal/common"
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

func (b *baseRepository[T]) GetList(model domain.BaseWhereModel, pagination *common.Pagination, orderBy *common.OrderBy) ([]T, error) {
	var ret []T
	db := model.SetExpression()(b.external.DB())
	if pagination != nil {
		db = db.Limit(pagination.Limit).Offset(pagination.Offset)
	}
	if orderBy != nil {
		db = db.Order(orderBy.ToClauseOrderBy())
	}

	err := db.Find(&ret).Error
	return ret, err
}

func newBaseRepository[T domain.ModelWithID](external domain.External) domain.BaseRepository[T] {
	err := external.DB().AutoMigrate(new(T))
	if err != nil {
		panic(err)
	}

	return &baseRepository[T]{
		external: external,
	}
}
