package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Nickname string `form:"nickname"`
	Email string `form:"email" notnull:"邮箱不能为空"`
	Password string `form:"password" notnull:"密码不能为空"`
	Salt int
	Ip string
}

type RegisterUser struct {
	Nickname string `form:"nickname" notnull:"昵称不能为空"`
	Email string `form:"email" notnull:"邮箱不能为空"`
	Password string `form:"password" notnull:"密码不能为空"`
	ConfirmPassword string `form:"confirm_password" notnull:"确认不能为空"`
}

func (User) TableName() string  {
	return "user"
}
