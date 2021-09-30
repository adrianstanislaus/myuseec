package responses

import (
	"myuseek/business/users"
	"time"
)

type UserRegisterResponse struct {
	Id                int       `json:"id"`
	FirstName         string    `json:"firstname"`
	LastName          string    `json:"lastname"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	Bio               string    `json:"bio"`
	Profile_pic       string    `json:"profile_pic"`
	Subscription_type string    `json:"subscription_type"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

func FromDomainToRegister(domain users.Domain) UserRegisterResponse {
	return UserRegisterResponse{
		Id:                domain.Id,
		FirstName:         domain.FirstName,
		LastName:          domain.LastName,
		Username:          domain.Username,
		Email:             domain.Email,
		Bio:               domain.Bio,
		Profile_pic:       domain.Profile_pic,
		Subscription_type: domain.Subscription_type,
		CreatedAt:         domain.CreatedAt,
		UpdatedAt:         domain.UpdatedAt,
	}
}
