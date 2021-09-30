package requests

import (
	"myuseek/business/users"
)

type UserByID struct {
	Id int `json:"id"`
}

func (user *UserByID) ToDomain() users.Domain {
	return users.Domain{
		Id: user.Id,
	}
}
