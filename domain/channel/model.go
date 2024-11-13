package channel

import "gorm.io/gorm"

type Channel struct {
	gorm.Model

	Name        string `gorm:"type:varchar(100);not null;unique"` // 채널 이름
	Description string `gorm:"type:text"`                         // 채널 설명
	IsPrivate   bool   `gorm:"default:false"`                     // 비공개 채널 여부
}
