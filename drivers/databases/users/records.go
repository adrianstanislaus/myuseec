package users

import (
	"myuseek/business/users"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id                int `gorm:"primaryKey"`
	FirstName         string
	LastName          string
	Username          string `gorm:"unique"`
	Email             string `gorm:"unique"`
	Password          string
	Bio               string
	Profile_pic       string
	Subscription_type string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

func (user *Users) ToDomain() users.Domain {
	return users.Domain{
		Id:                user.Id,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		Username:          user.Username,
		Email:             user.Email,
		Password:          user.Password,
		Bio:               user.Bio,
		Profile_pic:       user.Profile_pic,
		Subscription_type: user.Subscription_type,
		CreatedAt:         user.CreatedAt,
		UpdatedAt:         user.UpdatedAt,
	}
}

func FromDomain(domain users.Domain) Users {
	return Users{
		Id:                domain.Id,
		FirstName:         domain.FirstName,
		LastName:          domain.LastName,
		Username:          domain.Username,
		Email:             domain.Email,
		Password:          domain.Password,
		Bio:               domain.Bio,
		Profile_pic:       domain.Profile_pic,
		Subscription_type: domain.Subscription_type,
		CreatedAt:         domain.CreatedAt,
		UpdatedAt:         domain.UpdatedAt}
}

func ToListDomain(data []Users) (result []users.Domain) {
	result = []users.Domain{}
	for _, user := range data {
		result = append(result, user.ToDomain())
	}
	return result
}
