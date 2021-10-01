package requests

import "myuseek/business/songs"

type SongByID struct {
	Id int `json:"song_id"`
}

func (song *SongByID) ToDomain() songs.Domain {
	return songs.Domain{
		Id: song.Id,
	}
}
