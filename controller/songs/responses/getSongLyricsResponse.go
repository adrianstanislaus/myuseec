package responses

import (
	"myuseek/business/songs"
	"myuseek/controller/artists/responses"
)

type GetSongLyricsResponse struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Artist responses.GetArtistsResponse
	Lyrics string `json:"lyrics"`
}

func FromDomainToGetSongLyrics(domain songs.Domain) GetSongLyricsResponse {
	return GetSongLyricsResponse{
		Id:     domain.Id,
		Title:  domain.Title,
		Artist: responses.FromDomainToGetArtists(domain.Artist),
		Lyrics: domain.Lyrics,
	}
}
