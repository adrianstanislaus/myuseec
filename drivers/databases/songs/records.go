package songs

import (
	"myuseek/business/artists"
	"myuseek/business/songs"
	"time"

	"gorm.io/gorm"
)

type Song struct {
	Id        int    `gorm:"primaryKey"`
	Title     string `gorm:"uniqueIndex:idx_song;size:255"`
	Artist_id int    `gorm:"uniqueIndex:idx_song;type:BIGINT(255)"`
	Artist    Artist
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
		Artist:    song.Artist.ToDomain(),
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
		Artist:    FromArtistDomain(domain.Artist),
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

type Artist struct {
	Id           int
	Name         string
	About        string
	Record_label string
	Profile_pic  string
	Songs        []Song
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

func (artist *Artist) ToDomain() artists.Domain {
	return artists.Domain{
		Id:           artist.Id,
		Name:         artist.Name,
		About:        artist.About,
		Record_label: artist.Record_label,
		Profile_pic:  artist.Profile_pic,
		CreatedAt:    artist.CreatedAt,
		UpdatedAt:    artist.UpdatedAt,
	}
}

func FromArtistDomain(domain artists.Domain) Artist {
	return Artist{
		Id:           domain.Id,
		Name:         domain.Name,
		About:        domain.About,
		Record_label: domain.Record_label,
		Profile_pic:  domain.Profile_pic,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt}
}
