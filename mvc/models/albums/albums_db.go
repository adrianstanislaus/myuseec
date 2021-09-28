package albums

import (
	"time"

	"gorm.io/gorm"
)

type Album struct {
	Id           int            `gorm:"primaryKey" json:"id"`
	Title        string         `json:"title"`
	Album_type   string         `json:"album_type"`
	About        string         `json:"about"`
	Album_art    string         `json:"album_art"`
	Release_date string         `json:"release_date"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
