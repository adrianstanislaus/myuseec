package requests

import "myuseek/business/playlists"

type PlaylistCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Creator_id  int    `json:"creator_id"`
}

func (playlist *PlaylistCreate) ToDomain() playlists.Domain {
	return playlists.Domain{
		Name:        playlist.Name,
		Description: playlist.Description,
		Creator_id:  playlist.Creator_id,
	}
}
