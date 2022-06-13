package dao

import (
	db "ticket-backend/database"
	"ticket-backend/model"
)

func IsPlaneNumExist(name string) bool {
	var plane model.Plane
	db.DB.Where("plane_num=?", name).First(&plane)
	if plane.ID != 0 {
		return true
	}
	return false
}

func GetPlaneById(id uint) model.Plane {
	var plane model.Plane
	db.DB.Where("id=?", id).First(&plane)
	return plane
}

func GetPlaneList() []model.Plane {
	var plane []model.Plane
	db.DB.Find(&plane)
	return plane
}

func DeletePlaneById(id uint) model.Plane {
	var plane model.Plane
	db.DB.Where("id = ?", id).Delete(&plane)
}
