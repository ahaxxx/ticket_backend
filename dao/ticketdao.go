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
