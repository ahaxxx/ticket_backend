package dto

import "ticket-backend/model"

type AdminDto struct {
	Id     uint   `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Phone  string `json:"phone,omitempty"`
	Auth   uint   `json:"auth,omitempty"`
	Cid    uint   `json:"company,omitempty"`
	Status uint   `json:"status,omitempty"`
}

func ToAdminDto(admin model.Admin) AdminDto {
	return AdminDto{
		Id:     admin.ID,
		Name:   admin.Name,
		Phone:  admin.Phone,
		Auth:   admin.Auth,
		Cid:    admin.Cid,
		Status: admin.Status,
	}
}
