package message

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model

	ChannelID          uint   `gorm:"not null"`           // 채널 ID (foreign key)
	UserID             uint   `gorm:"not null"`           // 사용자 ID (foreign key)
	Content            string `gorm:"type:text;not null"` // 메시지 내용
	TranslatedContents string `gorm:"type:text;"`         // 메시지 내용
}

func (m Message) GetID() uint {
	return m.ID
}
