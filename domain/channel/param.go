package channel

import (
	"gorm.io/gorm"
	"uv-chat-api-server-golang/internal/common"
)

type ChannelParam struct {
	ID uint
}

func (p *ChannelParam) SetExpression() func(tx *gorm.DB) *gorm.DB {
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

type ChannelUsersParam struct {
	UserIDs   []int
	ChannelID uint
}

func (p *ChannelUsersParam) SetExpression() func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if len(p.UserIDs) > 0 {
			tx = tx.Where("user_id in (?)", p.UserIDs)
		}
		if p.ChannelID != 0 {
			tx = tx.Where("channel_id = ?", p.ChannelID)
		}
		return tx
	}
}
