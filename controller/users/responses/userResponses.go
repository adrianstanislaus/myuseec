package responses

import (
	"myuseek/business/users"
)

type GetUserByIdResponse struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Bio         string `json:"bio"`
	Profile_pic string `json:"profile_pic"`
}

func FromDomainToGetUserById(domain users.Domain) GetUserByIdResponse {
	return GetUserByIdResponse{
		Id:          domain.Id,
		FirstName:   domain.FirstName,
		LastName:    domain.LastName,
		Username:    domain.Username,
		Email:       domain.Email,
		Bio:         domain.Bio,
		Profile_pic: domain.Profile_pic,
	}
}
