package channel

import "gorm.io/gorm"

type ChannelParam struct {
}

func (p *ChannelParam) SetExpression() func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx
	}
}
