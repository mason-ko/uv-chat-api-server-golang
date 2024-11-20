package user

import (
	"gorm.io/gorm"
)

type User struct {
	ID uint

	Name    string
	Country string
}

func (m User) DBModel() DBUser {
	return DBUser{
		Name:    m.Name,
		Country: m.Country,
	}
}

type DBUser struct {
	gorm.Model

	Name    string `gorm:"type:varchar(100);"`
	Country string `gorm:"type:varchar(100);"`
}

func (m DBUser) GetID() uint {
	return m.ID
}

func (m DBUser) User() User {
	return User{
		ID:      m.ID,
		Name:    m.Name,
		Country: m.Country,
	}
}
