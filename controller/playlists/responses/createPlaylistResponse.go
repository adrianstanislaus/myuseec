package responses

import (
	"myuseek/business/playlists"
)

type CreatePlaylistResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Creator_id  int    `json:"creator_id"`
}

func FromDomainToCreatePlaylist(domain playlists.Domain) CreatePlaylistResponse {
	return CreatePlaylistResponse{
		Name:        domain.Name,
		Description: domain.Description,
		Creator_id:  domain.Creator_id,
	}
}
