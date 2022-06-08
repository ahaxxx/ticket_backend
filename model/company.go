package model

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name string `gorm:"type:varchar(25);not null",json:"name"`
}
