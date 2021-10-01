package responses

import (
	"myuseek/business/songs"
)

type SongsInPlaylistResponse struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Artist_id int    `json:"artist_id"`
	Album_id  int    `json:"album_id"`
	Genre     string `json:"genre"`
	Duration  string `json:"duration"`
}

func FromDomainToPlaylistSongs(domain songs.Domain) SongsInPlaylistResponse {
	return SongsInPlaylistResponse{
		Id:        domain.Id,
		Title:     domain.Title,
		Artist_id: domain.Artist_id,
		Album_id:  domain.Album_id,
		Genre:     domain.Genre,
		Duration:  domain.Duration}
}

func FromListDomainToPlaylistSongs(data []songs.Domain) (result []SongsInPlaylistResponse) {
	result = []SongsInPlaylistResponse{}
	for _, song := range data {
		result = append(result, FromDomainToPlaylistSongs(song))
	}
	return result
}
