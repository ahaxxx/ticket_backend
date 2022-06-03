package dao

import (
	db "ticket-backend/database"
	"ticket-backend/model"
)

func GetPassengerList(uid uint) []model.Passenger {
	var passenger []model.Passenger
	db.DB.Where("uid=?", uid).Find(&passenger)
	return passenger
}
