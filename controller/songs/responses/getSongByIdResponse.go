package responses

import (
	"myuseek/business/songs"
	"myuseek/controller/artists/responses"
)

type GetSongByIdResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Artist   responses.GetArtistsResponse
	Album_id int    `json:"album_id"`
	Genre    string `json:"genre"`
	Duration string `json:"duration"`
	Lyrics   string `json:"lyrics"`
}

func FromDomainToGetSongById(domain songs.Domain) GetSongByIdResponse {
	return GetSongByIdResponse{
		Id:       domain.Id,
		Title:    domain.Title,
		Artist:   responses.FromDomainToGetArtists(domain.Artist),
		Album_id: domain.Album_id,
		Genre:    domain.Genre,
		Duration: domain.Duration,
		Lyrics:   domain.Lyrics,
	}
}
