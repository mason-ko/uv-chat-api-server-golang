package message

import (
	"gorm.io/gorm"
	"uv-chat-api-server-golang/internal/common"
)

type DBMessageParam struct {
	ID        uint
	ChannelID uint `json:"channel_id"`
}

func (p *DBMessageParam) SetExpression() func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if p.ID != 0 {
			tx = tx.Where("id = ?", p.ID)
		}

		if p.ChannelID != 0 {
			tx = tx.Where("channel_id = ?", p.ChannelID)
		}
		return tx
	}
}

type GetListParam struct {
	Pagination *common.Pagination
	OrderBy    *common.OrderBy
}
