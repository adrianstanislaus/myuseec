package responses

import (
	"myuseek/business/songs"
	"myuseek/controller/artists/responses"
)

type GetSongsResponse struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Artist   responses.GetArtistsResponse
	Album_id int    `json:"album_id"`
	Genre    string `json:"genre"`
	Duration string `json:"duration"`
}

func FromDomainToGetSongs(domain songs.Domain) GetSongsResponse {
	return GetSongsResponse{
		Id:       domain.Id,
		Title:    domain.Title,
		Artist:   responses.FromDomainToGetArtists(domain.Artist),
		Album_id: domain.Album_id,
		Genre:    domain.Genre,
		Duration: domain.Duration}
}

func FromListDomainToGetSongs(data []songs.Domain) (result []SongResponse) {
	result = []SongResponse{}
	for _, song := range data {
		result = append(result, FromDomain(song))
	}
	return result
}
