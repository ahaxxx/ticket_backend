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
