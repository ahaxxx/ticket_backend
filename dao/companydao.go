package dao

import (
	db "ticket-backend/database"
	"ticket-backend/model"
)

func GetCompanyList() []model.Company {
	var company []model.Company
	db.DB.Find(&company)
	return company
}

func IsCompanyExist(name string) bool {
	var company model.Company
	db.DB.Where("name=?", name).First(&name)
	if company.ID != 0 {
		return true
	}
	return false
}
