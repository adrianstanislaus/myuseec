package requests

import (
	"myuseek/business/albums"
)

type AlbumByID struct {
	Id int `json:"album_id"`
}

func (album *AlbumByID) ToDomain() albums.Domain {
	return albums.Domain{
		Id: album.Id,
	}
}
