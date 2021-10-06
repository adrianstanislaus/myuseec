package requests

import "myuseek/business/playlists"

type PlaylistSearch struct {
	Name string `json:"id"`
}

func (playlist *PlaylistSearch) ToDomain() playlists.Domain {
	return playlists.Domain{
		Name: playlist.Name,
	}
}
