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

func DeletePlaneById(id uint) {
	var plane model.Plane
	db.DB.Where("id = ?", id).Delete(&plane)
}

func SearchPlaneByDAT(departure string, arrival string, takeoffTime uint) []model.Plane {
	var planes []model.Plane
	db.DB.Where(db.DB.Where("departure = ?", departure).Where(db.DB.Where("arrival = ?", arrival)).Where(db.DB.Where("takeoff_time = ?", takeoffTime))).Find(&model.Plane{})
	return planes
}

func UpdatePlaneById(id uint, plane model.Plane) {
	db.DB.Model(&model.Plane{}).Where("id=?", id).Update(&plane)
}
