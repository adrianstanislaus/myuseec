package responses

import (
	"myuseek/business/users"
)

type GetUsersResponse struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Profile_pic string `json:"profile_pic"`
}

func FromDomainToGetUsers(domain users.Domain) GetUsersResponse {
	return GetUsersResponse{
		Id:          domain.Id,
		Username:    domain.Username,
		Email:       domain.Email,
		Profile_pic: domain.Profile_pic,
	}
}

func FromListDomainToGetUsers(data []users.Domain) (result []GetUsersResponse) {
	result = []GetUsersResponse{}
	for _, user := range data {
		result = append(result, FromDomainToGetUsers(user))
	}
	return result
}
