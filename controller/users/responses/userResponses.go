package responses

import (
	"myuseek/business/users"
	"time"
)

type UserResponse struct {
	Id                int       `json:"id"`
	FirstName         string    `json:"firstname"`
	LastName          string    `json:"lastname"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	Bio               string    `json:"bio"`
	Profile_pic       string    `json:"profile_pic"`
	Subscription_type string    `json:"subscription_type"`
	Token             string    `json:"token"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		Id:                domain.Id,
		FirstName:         domain.FirstName,
		LastName:          domain.LastName,
		Username:          domain.Username,
		Email:             domain.Email,
		Password:          domain.Password,
		Bio:               domain.Bio,
		Profile_pic:       domain.Profile_pic,
		Subscription_type: domain.Subscription_type,
		Token:             domain.Token,
		CreatedAt:         domain.CreatedAt,
		UpdatedAt:         domain.UpdatedAt,
	}
}
