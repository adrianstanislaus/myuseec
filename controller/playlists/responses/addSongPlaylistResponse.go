package responses

import (
	"myuseek/business/playlists"
	"myuseek/controller/songs/responses"
)

type AddSongPlaylistResponse struct {
	Name       string                              `json:"name"`
	Creator_id int                                 `json:"creator_id"`
	Songs      []responses.SongsInPlaylistResponse `json:"added_song"`
}

func FromDomainToAddSong(domain playlists.Domain) AddSongPlaylistResponse {
	return AddSongPlaylistResponse{
		Name:       domain.Name,
		Creator_id: domain.Creator_id,
		Songs:      responses.FromListDomainToPlaylistSongs(domain.Songs),
	}
}
