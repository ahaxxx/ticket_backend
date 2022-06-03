package model

import "gorm.io/gorm"

type Passenger struct {
	gorm.Model
	Name     string `gorm:"type:varchar(25);not null",json:"name"`
	Idnum    string `gorm:"type:varchar(255);not null",json:"idnum"`
	Phone    string `gorm:"type:varchar(255);not null",json:"phone"`
	Sex      string `gorm:"type:varchar(255);not null",json:"sex"`
	Workunit string `gorm:"type:varchar(255);not null",json:"workunit"`
	UID      uint
}
