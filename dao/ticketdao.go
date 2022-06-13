package dao

import (
	db "ticket-backend/database"
	"ticket-backend/model"
)

func GetTicketListByUid(uid uint) []model.Ticket {
	var tickets []model.Ticket
	db.DB.Where("uid=?", uid).Find(&tickets)
	return tickets
}

func GetTicketById(id uint) model.Ticket {
	var ticket model.Ticket
	db.DB.Where("id=?", id).First(&ticket)
	return ticket
}

func UpdateTicketById(id uint, ticket model.Ticket) {
	db.DB.Model(&model.Ticket{}).Where("id=?", id).Update(&ticket)
}
