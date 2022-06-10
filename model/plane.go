package model

import (
	"gorm.io/gorm"
)

type Plane struct {
	gorm.Model
	PlaneNum    string  `gorm:"type:varchar(25);not null",json:"planenum"`
	Status      string  `gorm:"type:varchar(255);not null",json:"status"`
	Price       float32 `gorm:"not null",json:"price"`
	Seat        int     `gorm:"not null",json:"seat"`
	Departure   string  `gorm:"type:varchar(255);not null",json:"departure"`
	Arrival     string  `gorm:"type:varchar(255);not null",json:"arrival"`
	TakeoffTime uint    `gorm:"not null"`
	Company     Company `json:"company"`
	CompanyId   int     `json:"company_id"`
}
