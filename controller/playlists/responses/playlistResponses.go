package responses

import (
	"myuseek/business/playlists"
	"myuseek/business/songs"
)

type PlaylistResponse struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Creator_id  int            `json:"creator_id"`
	Songs       []songs.Domain `json:"songs_added"`
}

func FromDomain(domain playlists.Domain) PlaylistResponse {
	return PlaylistResponse{
		Name:        domain.Name,
		Description: domain.Description,
		Creator_id:  domain.Creator_id,
		Songs:       domain.Songs,
	}
}

func FromListDomain(data []playlists.Domain) (result []PlaylistResponse) {
	result = []PlaylistResponse{}
	for _, playlist := range data {
		result = append(result, FromDomain(playlist))
	}
	return result
}
