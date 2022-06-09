package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string `gorm:"type:varchar(25);not null",json:"name"`
	Auth     uint   `gorm:"not null",json:"Auth"`
	Password string `gorm:"type:varchar(255);not null",json:"password"`
	Phone    string `gorm:"type:varchar(255);not null",json:"phone"`
	Cid      uint   `gorm:"not null",json:"cid"`
	Status   uint   `gorm:"not null",json:"status"`
}
