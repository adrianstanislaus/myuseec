package responses

import (
	"myuseek/business/playlists"
	"myuseek/controller/songs/responses"
)

type RemoveSongPlaylistResponse struct {
	Name       string                              `json:"name"`
	Creator_id int                                 `json:"creator_id"`
	Songs      []responses.SongsInPlaylistResponse `json:"removed_song"`
}

func FromDomainToRemoveSong(domain playlists.Domain) RemoveSongPlaylistResponse {
	return RemoveSongPlaylistResponse{
		Name:       domain.Name,
		Creator_id: domain.Creator_id,
		Songs:      responses.FromListDomainToPlaylistSongs(domain.Songs),
	}
}
