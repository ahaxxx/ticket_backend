package model

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Uid     uint    `gorm:"not null",json:"uid"`
	PlaneId uint    `gorm:"not null",json:"plane_id"`
	PassId  uint    `gorm:"not null"'json:"pass_id"`
	Price   float32 `gorm:"not null",json:"price"`
	Status  uint    `gorm:"not null",json:"status"`
}
