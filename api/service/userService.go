package service

import "newbug/model"

func Store(user *model.User) int64 {
	return model.Connect().Create(&user).RowsAffected
}

func FindByEmail(user *model.User) {
	model.Connect().Where("email = ?", "sdf").First(user)
}
