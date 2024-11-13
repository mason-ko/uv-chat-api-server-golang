package message

import (
	"gorm.io/gorm"
	"uv-chat-api-server-golang/internal/common"
)

type MessageParam struct {
	ChannelID uint `json:"channel_id"`
}

func (p *MessageParam) SetExpression() func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
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
