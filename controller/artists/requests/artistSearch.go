package requests

import (
	"myuseek/business/artists"
)

type ArtistSearch struct {
	Name string `json:"name"`
}

func (artist *ArtistSearch) ToDomain() artists.Domain {
	return artists.Domain{
		Name: artist.Name,
	}
}
