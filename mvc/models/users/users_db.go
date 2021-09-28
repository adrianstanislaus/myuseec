package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id                int            `gorm:"primaryKey" json:"id"`
	FirstName         string         `json:"firstname"`
	LastName          string         `json:"lastname"`
	Username          string         `json:"username" gorm:"unique"`
	Email             string         `json:"email" gorm:"unique"`
	Password          string         `json:"password"`
	Bio               string         `json:"bio"`
	Profile_pic       string         `json:"profile_pic"`
	Subscription_type string         `json:"subscription_type"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
