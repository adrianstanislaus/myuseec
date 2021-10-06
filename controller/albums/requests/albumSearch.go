package requests

import (
	"myuseek/business/albums"
)

type AlbumSearch struct {
	Title string `json:"album_id"`
}

func (album *AlbumSearch) ToDomain() albums.Domain {
	return albums.Domain{
		Title: album.Title,
	}
}
