package responses

import (
	"myuseek/business/albums"
	"myuseek/controller/songs/responses"
	"time"
)

type AlbumResponse struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Artist_id    int    `json:"artist_id"`
	Songs        []responses.SongsInAlbumResponse
	Album_art    string    `json:"album_art"`
	Release_date time.Time `json:"release_date"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func FromDomain(domain albums.Domain) AlbumResponse {
	return AlbumResponse{
		Id:           domain.Id,
		Title:        domain.Title,
		Artist_id:    domain.Artist_id,
		Songs:        responses.FromListDomainToAlbumSongs(domain.Songs),
		Album_art:    domain.Album_art,
		Release_date: domain.Release_date,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt}
}

func FromListDomain(data []albums.Domain) (result []AlbumResponse) {
	result = []AlbumResponse{}
	for _, album := range data {
		result = append(result, FromDomain(album))
	}
	return result
}
