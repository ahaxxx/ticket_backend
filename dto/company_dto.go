package dto

import "ticket-backend/model"

type CompanyDto struct {
	Id   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func ToCompanyDto(company model.Company) CompanyDto {
	return CompanyDto{
		Id:   company.ID,
		Name: company.Name,
	}
}
