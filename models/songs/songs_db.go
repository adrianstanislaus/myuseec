package Songs

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id           int            `gorm:"primaryKey" json:"id"`
	Title        string         `json:"title"`
	Artist_id    string         `json:"artist_id"`
	Album_id     string         `json:"album_id"`
	Record_label string         `json:"record_label"`
	Duration     string         `json:"duration"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
