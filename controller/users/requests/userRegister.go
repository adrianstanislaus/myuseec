package requests

import "myuseek/business/users"

type UserRegister struct {
	FirstName         string `json:"firstname"`
	LastName          string `json:"lastname"`
	Username          string `json:"username"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	Bio               string `json:"bio"`
	Profile_pic       string `json:"profile_pic"`
	Subscription_type string `json:"subscription_type"`
}

func (user *UserRegister) ToDomain() users.Domain {
	return users.Domain{
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		Username:          user.Username,
		Email:             user.Email,
		Password:          user.Password,
		Bio:               user.Bio,
		Profile_pic:       user.Profile_pic,
		Subscription_type: user.Subscription_type,
	}
}
