package requests

import "myuseek/business/songs"

type SongAdd struct {
	Title     string `json:"title"`
	Artist_id int    `json:"artist_id"`
	Album_id  int    `json:"album_id"`
	Genre     string `json:"genre"`
	Duration  string `json:"duration"`
	Lyrics    string `json:"lyrics"`
}

func (song *SongAdd) ToDomain() songs.Domain {
	return songs.Domain{
		Title:     song.Title,
		Artist_id: song.Artist_id,
		Album_id:  song.Album_id,
		Genre:     song.Genre,
		Duration:  song.Duration,
		Lyrics:    song.Lyrics,
	}
}
