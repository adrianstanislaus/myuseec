package responses

import (
	"myuseek/business/albums"
	"myuseek/controller/songs/responses"
	"time"
)

type GetAlbumByIdResponse struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Artist_id    int    `json:"artist_id"`
	Songs        []responses.SongsInAlbumResponse
	Album_art    string    `json:"album_art"`
	Release_date time.Time `json:"release_date"`
}

func FromDomainToGetAlbumById(domain albums.Domain) GetAlbumByIdResponse {
	return GetAlbumByIdResponse{
		Id:           domain.Id,
		Title:        domain.Title,
		Artist_id:    domain.Artist_id,
		Songs:        responses.FromListDomainToAlbumSongs(domain.Songs),
		Album_art:    domain.Album_art,
		Release_date: domain.Release_date,
	}
}
