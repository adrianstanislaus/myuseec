package playlists

import (
	"time"

	"gorm.io/gorm"
)

type Playlists struct {
	Id          int            `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"descrition"`
	Creator_id  string         `json:"creator_id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
