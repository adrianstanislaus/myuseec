package lyricsapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Lyricsapi struct {
	Lyrics string `json:"lyrics"`
}

func GetLyrics(title, artist string) (string, error) {
	apilink := "https://api.lyrics.ovh/v1/" + artist + "/" + title
	response, err_api := http.Get(apilink)

	if err_api != nil {
		var lyrics string
		return lyrics, err_api
	}

	responseData, err_read := ioutil.ReadAll(response.Body)

	if err_read != nil {
		var lyrics string
		return lyrics, err_read
	}

	defer response.Body.Close()
	var lyrics Lyricsapi
	json.Unmarshal(responseData, &lyrics)
	return lyrics.Lyrics, nil
}
