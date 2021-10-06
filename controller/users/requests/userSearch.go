package requests

import (
	"myuseek/business/users"
)

type UserSearch struct {
	Username string `json:"title"`
}

func (user *UserSearch) ToDomain() users.Domain {
	return users.Domain{
		Username: user.Username,
	}
}
