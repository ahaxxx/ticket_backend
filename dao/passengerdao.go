package dao

import (
	db "ticket-backend/database"
	"ticket-backend/model"
)

func GetPassengerListByUid(uid uint) []model.Passenger {
	var passenger []model.Passenger
	db.DB.Where("uid=?", uid).Find(&passenger)
	return passenger
}

func GetPassengerById(id uint) model.Passenger {
	var passenger model.Passenger
	db.DB.Where("id=?", id).First(&passenger)
	return passenger
}

func UpdatePassengerById(id uint, passenger model.Passenger) {
	db.DB.Model(&model.Passenger{}).Where("id=?", id).Update(&passenger)
}

func DeletePassengerById(id uint) {
	var passenger model.Passenger
	db.DB.Where("id = ?", id).Delete(&passenger)
}
