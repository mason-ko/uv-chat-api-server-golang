package channel

import "gorm.io/gorm"

type Channel struct {
	ID uint `json:"id"`

	Name        string `json:"name" gorm:"type:(m *messageController)(100);not null;unique"` // 채널 이름
	Description string `json:"description" gorm:"type:text"`                                 // 채널 설명
	IsPrivate   bool   `json:"isPrivate" gorm:"default:false"`                               // 비공개 채널 여부
	LastContent string `json:"lastContent" gorm:"type:text"`                                 // 마지막 content
}

func (m Channel) DBModel() DBChannel {
	return DBChannel{
		Name:        m.Name,
		Description: m.Description,
		IsPrivate:   m.IsPrivate,
		LastContent: m.LastContent,
	}
}

type DBChannel struct {
	gorm.Model

	Name        string `gorm:"type:varchar(100);not null;unique"` // 채널 이름
	Description string `gorm:"type:text"`                         // 채널 설명
	IsPrivate   bool   `gorm:"default:false"`                     // 비공개 채널 여부
	LastContent string `gorm:"type:text"`                         // 마지막 content
}

func (m DBChannel) GetID() uint {
	return m.ID
}

func (m DBChannel) Channel() Channel {
	return Channel{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		IsPrivate:   m.IsPrivate,
		LastContent: m.LastContent,
	}
}

type DBChannelUsers struct {
	gorm.Model

	ChannelID uint `gorm:"index:idx_channel_user,unique"`
	UserID    uint `gorm:"index:idx_channel_user,unique"`
}

func (m DBChannelUsers) GetID() uint {
	return m.ID
}

type ReqCreateChannel struct {
	UserID       int `json:"userId" binding:"required"`
	TargetUserID int `json:"targetUserId" binding:"required"`
}
