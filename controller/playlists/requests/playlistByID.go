package requests

import "myuseek/business/playlists"

type PlaylistByID struct {
	Id int `json:"id"`
}

func (playlist *PlaylistByID) ToDomain() playlists.Domain {
	return playlists.Domain{
		Id: playlist.Id,
	}
}
