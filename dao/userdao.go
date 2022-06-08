package dao

import (
	db "ticket-backend/database"
	"ticket-backend/model"
)

//
//  IsPhoneExist
//  @Description:判断手机号是否已存在
//  @param phone
//  @return bool
//
func IsUserPhoneExist(phone string) bool {
	var user model.User
	db.DB.Where("phone=?", phone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

//
//  IsNameExist
//  @Description:判断用户名是否存在
//  @param name
//  @return bool
//
func IsUserNameExist(name string) bool {
	var user model.User
	db.DB.Where("name=?", name).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func UpdateUserById(id uint, user model.User) {
	db.DB.Model(&model.User{}).Where("id=?", id).Update(&user)
}
