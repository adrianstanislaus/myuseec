package requests

import (
	"myuseek/business/playlists"
	"myuseek/controller/songs/requests"
)

type SongChoosePlaylist struct {
	Id         int               `json:"playlist_id"`
	SongChoose requests.SongByID `json:"song"`
}

func (playlist *SongChoosePlaylist) ToDomain() playlists.Domain {
	var playlistDomain playlists.Domain
	playlistDomain.Id = playlist.Id
	playlistDomain.Songs = append(playlistDomain.Songs, playlist.SongChoose.ToDomain())
	return playlistDomain
}
