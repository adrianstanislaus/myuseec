package artists

import (
	"time"

	"gorm.io/gorm"
)

type Artist struct {
	Id          int            `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name"`
	About       string         `json:"about"`
	Profile_pic string         `json:"profile_pic"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
