package requests

import (
	"myuseek/business/albums"
	"time"
)

type AlbumAdd struct {
	Title        string    `json:"title"`
	Artist_id    int       `json:"artist_id"`
	Album_art    string    `json:"album_art"`
	Release_date time.Time `json:"release_date"`
}

func (album *AlbumAdd) ToDomain() albums.Domain {
	return albums.Domain{
		Title:        album.Title,
		Artist_id:    album.Artist_id,
		Album_art:    album.Album_art,
		Release_date: album.Release_date,
	}
}
