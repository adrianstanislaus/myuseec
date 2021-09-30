package responses

import (
	"myuseek/business/playlists"
)

type GetPlaylistsResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Creator_id  int    `json:"creator_id"`
}

func FromDomainToGetPlaylists(domain playlists.Domain) GetPlaylistsResponse {
	return GetPlaylistsResponse{
		Name:        domain.Name,
		Description: domain.Description,
		Creator_id:  domain.Creator_id,
	}
}

func FromListDomainToGetPlaylists(data []playlists.Domain) (result []GetPlaylistsResponse) {
	result = []GetPlaylistsResponse{}
	for _, playlist := range data {
		result = append(result, FromDomainToGetPlaylists(playlist))
	}
	return result
}
