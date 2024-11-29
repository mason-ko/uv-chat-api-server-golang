package message

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID                uint   `json:"id"`
	ChannelID         uint   `json:"channelId"`
	UserID            uint   `json:"userId"`
	Content           string `json:"content"`
	TranslatedContent string `json:"translatedContent"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (m Message) DBModel() DBMessage {
	return DBMessage{
		ChannelID:         m.ChannelID,
		UserID:            m.UserID,
		Content:           m.Content,
		TranslatedContent: m.TranslatedContent,
	}
}

type DBMessage struct {
	gorm.Model

	ChannelID         uint   `gorm:"not null"`           // 채널 ID (foreign key)
	UserID            uint   `gorm:"not null"`           // 사용자 ID (foreign key)
	Content           string `gorm:"type:text;not null"` // 메시지 내용
	TranslatedContent string `gorm:"type:text;"`         // 메시지 내용
}

func (m DBMessage) GetID() uint {
	return m.ID
}

func (m DBMessage) Message() Message {
	return Message{
		ID:                m.ID,
		ChannelID:         m.ChannelID,
		UserID:            m.UserID,
		Content:           m.Content,
		TranslatedContent: m.TranslatedContent,
		CreatedAt:         m.CreatedAt,
		UpdatedAt:         m.UpdatedAt,
	}
}

type ReqCreateMessage struct {
	ChannelID uint   `json:"channelId"`
	Content   string `json:"content"`
	UserID    uint   `json:"userId"`
}
