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

func GetCompanyByCid(cid uint) model.Company {
	var company model.Company
	db.DB.Where("id=?", cid).Find(&company)
	return company
}

func UpdateCompanyById(cid uint, company model.Company) {
	db.DB.Model(&model.Company{}).Where("id=?", cid).Update(&company)
}

func DeleteCompanyById(cid uint) {
	var company model.Company
	db.DB.Where("id = ?", cid).Delete(&company)
}
