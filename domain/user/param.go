package user

import (
	"gorm.io/gorm"
	"uv-chat-api-server-golang/internal/common"
)

type DBUserParam struct {
	ID uint
}

func (p *DBUserParam) SetExpression() func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if p.ID != 0 {
			tx = tx.Where("id = ?", p.ID)
		}
		return tx
	}
}

type GetListParam struct {
	Pagination *common.Pagination
	OrderBy    *common.OrderBy
}
