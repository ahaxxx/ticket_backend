package dto

import "ticket-backend/model"

type UserDto struct {
	Id    uint   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Id:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
	}
}
