package requests

import "myuseek/business/songs"

type SongSearch struct {
	Title string `json:"title"`
}

func (song *SongSearch) ToDomain() songs.Domain {
	return songs.Domain{
		Title: song.Title,
	}
}
