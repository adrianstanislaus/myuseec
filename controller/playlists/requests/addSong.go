package requests

import (
	"myuseek/business/playlists"
	"myuseek/business/songs"
)

type NewSong struct {
	Id        int          `json:"song_id"`
	SongToAdd songs.Domain `json:"song"`
}

func (song *NewSong) ToDomain() playlists.Domain {
	var songDomain playlists.Domain
	songDomain.Id = song.Id
	songDomain.Songs = append(songDomain.Songs, song.SongToAdd)
	return songDomain
}
