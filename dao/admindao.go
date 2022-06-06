package dao

import (
	db "ticket-backend/database"
	"ticket-backend/model"
)

func IsAdminPhoneExist(phone string) bool {
	var admin model.Admin
	db.DB.Where("phone=?", phone).First(&admin)
	if admin.ID != 0 {
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
func IsAdminNameExist(name string) bool {
	var admin model.Admin
	db.DB.Where("name=?", name).First(&admin)
	if admin.ID != 0 {
		return true
	}
	return false
}
