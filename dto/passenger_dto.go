package dto

import "ticket-backend/model"

type PassengerDto struct {
	Id       uint   `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	IdNum    string `json:"id_num,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Sex      string `json:"sex,omitempty"`
	Workunit string `json:"workunit,omitempty"`
}

func ToPassengerDto(passenger model.Passenger) PassengerDto {
	return PassengerDto{
		Id:       passenger.ID,
		Name:     passenger.Name,
		IdNum:    passenger.Idnum,
		Phone:    passenger.Phone,
		Sex:      passenger.Sex,
		Workunit: passenger.Workunit,
	}
}
