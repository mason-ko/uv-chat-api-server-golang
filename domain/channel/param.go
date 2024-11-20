package channel

import (
	"gorm.io/gorm"
	"uv-chat-api-server-golang/internal/common"
)

type ChannelParam struct {
}

func (p *ChannelParam) SetExpression() func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx
	}
}

type GetListParam struct {
	Pagination *common.Pagination
	OrderBy    *common.OrderBy
}
