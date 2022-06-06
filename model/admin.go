package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string `gorm:"type:varchar(25);not null",json:"name"`
	Auth     uint
	Password string `gorm:"type:varchar(255);not null",json:"password"`
	Phone    string `gorm:"type:varchar(255);not null",json:"phone"`
	Company  string `gorm:"type:varchar(255);not null",json:"company"`
}
