package responses

import (
	"myuseek/business/users"
)

type UserLoginResponse struct {
	Id                int    `json:"id"`
	FirstName         string `json:"firstname"`
	LastName          string `json:"lastname"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	Subscription_type string `json:"subscription_type"`
	Token             string `json:"token"`
}

func FromDomainToLogin(domain users.Domain) UserLoginResponse {
	return UserLoginResponse{
		Id:                domain.Id,
		FirstName:         domain.FirstName,
		LastName:          domain.LastName,
		Username:          domain.Username,
		Email:             domain.Email,
		Subscription_type: domain.Subscription_type,
		Token:             domain.Token,
	}
}
