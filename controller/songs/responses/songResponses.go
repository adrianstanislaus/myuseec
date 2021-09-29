package responses

import (
	"myuseek/business/songs"
	"myuseek/controller/artists/responses"
	"time"
)

type SongResponse struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Artist_id int    `json:"artist_id"`
	Artist    responses.ArtistResponse
	Album_id  int       `json:"album_id"`
	Genre     string    `json:"genre"`
	Duration  string    `json:"duration"`
	Lyrics    string    `json:"lyrics"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FromDomain(domain songs.Domain) SongResponse {
	return SongResponse{
		Id:        domain.Id,
		Title:     domain.Title,
		Artist_id: domain.Artist_id,
		Artist:    responses.FromDomain(domain.Artist),
		Album_id:  domain.Album_id,
		Genre:     domain.Genre,
		Duration:  domain.Duration,
		Lyrics:    domain.Lyrics,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt}
}

func FromListDomain(data []songs.Domain) (result []SongResponse) {
	result = []SongResponse{}
	for _, song := range data {
		result = append(result, FromDomain(song))
	}
	return result
}
