package domain

import "gorm.io/gorm"

type External interface {
	DB() *gorm.DB
}
