package songs

import (
	"myuseek/business/songs"
	"time"

	"gorm.io/gorm"
)

type Song struct {
	Id        int    `gorm:"primaryKey"`
	Title     string `gorm:"uniqueIndex:idx_song;size:255"`
	Artist_id int    `gorm:"uniqueIndex:idx_song;type:BIGINT(255)"`
	Album_id  int
	Genre     string
	Duration  string
	Lyrics    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (song *Song) ToDomain() songs.Domain {
	return songs.Domain{
		Id:        song.Id,
		Title:     song.Title,
		Artist_id: song.Artist_id,
		Album_id:  song.Album_id,
		Genre:     song.Genre,
		Duration:  song.Duration,
		Lyrics:    song.Lyrics,
		CreatedAt: song.CreatedAt,
		UpdatedAt: song.UpdatedAt,
	}
}

func FromDomain(domain songs.Domain) Song {
	return Song{
		Id:        domain.Id,
		Title:     domain.Title,
		Artist_id: domain.Artist_id,
		Album_id:  domain.Album_id,
		Genre:     domain.Genre,
		Duration:  domain.Duration,
		Lyrics:    domain.Lyrics,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt}
}

func ToListDomain(data []Song) (result []songs.Domain) {
	result = []songs.Domain{}
	for _, song := range data {
		result = append(result, song.ToDomain())
	}
	return result
}

func FromListDomain(data []songs.Domain) (result []Song) {
	result = []Song{}
	for _, song := range data {
		result = append(result, FromDomain(song))
	}
	return result
}
