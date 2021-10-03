package responses

import (
	"myuseek/business/songs"
)

type SongsInAlbumResponse struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Artist_id int    `json:"artist_id"`
}

func FromDomainToAlbumSongs(domain songs.Domain) SongsInAlbumResponse {
	return SongsInAlbumResponse{
		Id:        domain.Id,
		Title:     domain.Title,
		Artist_id: domain.Artist_id,
	}
}

func FromListDomainToAlbumSongs(data []songs.Domain) (result []SongsInAlbumResponse) {
	result = []SongsInAlbumResponse{}
	for _, song := range data {
		result = append(result, FromDomainToAlbumSongs(song))
	}
	return result
}
