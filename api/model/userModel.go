package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Nickname string
	Email string
	Password string
	Salt int
	Ip string
}

func (User) TableName() string  {
	return "user"
}