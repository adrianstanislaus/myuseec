package albums

import (
	"myuseek/business/albums"
	"myuseek/drivers/databases/songs"
	"time"

	"gorm.io/gorm"
)

type Album struct {
	Id           int          `gorm:"primaryKey"`
	Title        string       `gorm:"uniqueIndex:idx_album;size:255"`
	Artist_id    int          `gorm:"uniqueIndex:idx_album;type:BIGINT(255)"`
	Songs        []songs.Song `gorm:"foreignKey:Album_id"`
	Album_art    string
	Release_date time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (album *Album) ToDomain() albums.Domain {
	return albums.Domain{
		Id:           album.Id,
		Title:        album.Title,
		Artist_id:    album.Artist_id,
		Songs:        songs.ToListDomain(album.Songs),
		Album_art:    album.Album_art,
		Release_date: album.Release_date,
		CreatedAt:    album.CreatedAt,
		UpdatedAt:    album.UpdatedAt,
	}
}

func FromDomain(domain albums.Domain) Album {
	return Album{
		Id:           domain.Id,
		Title:        domain.Title,
		Artist_id:    domain.Artist_id,
		Songs:        songs.FromListDomain(domain.Songs),
		Album_art:    domain.Album_art,
		Release_date: domain.Release_date,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt}
}

func ToListDomain(data []Album) (result []albums.Domain) {
	result = []albums.Domain{}
	for _, album := range data {
		result = append(result, album.ToDomain())
	}
	return result
}
