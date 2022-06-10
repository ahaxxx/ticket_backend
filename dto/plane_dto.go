package dto

import "ticket-backend/model"

type PlaneDto struct {
	PlaneNum    string  `json:"plane_num"`
	Status      string  `json:"status"`
	Price       float32 `json:"price"`
	Seat        int     `json:"seat"`
	Departure   string  `json:"departure"`
	Arrival     string  `json:"arrival"`
	TakeoffTime uint    `json:"takeoff_time"`
	CompanyId   int     `json:"company_id"`
}

func ToPlaneDto(plane model.Plane) PlaneDto {
	return PlaneDto{
		PlaneNum:    plane.PlaneNum,
		Status:      plane.Status,
		Price:       plane.Price,
		Seat:        plane.Seat,
		Departure:   plane.Departure,
		Arrival:     plane.Arrival,
		TakeoffTime: plane.TakeoffTime,
		CompanyId:   plane.CompanyId,
	}
}
