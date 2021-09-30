package requests

import (
	"myuseek/business/artists"
)

type ArtistByID struct {
	Id int `json:"id"`
}

func (artist *ArtistByID) ToDomain() artists.Domain {
	return artists.Domain{
		Id: artist.Id,
	}
}
