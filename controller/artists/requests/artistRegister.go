package requests

import (
	"myuseek/business/artists"
)

type ArtistRegister struct {
	Name         string `json:"name"`
	About        string `json:"about"`
	Record_label string `json:"record_label"`
	Profile_pic  string `json:"profile_pic"`
}

func (artist *ArtistRegister) ToDomain() artists.Domain {
	return artists.Domain{
		Name:         artist.Name,
		About:        artist.About,
		Record_label: artist.Record_label,
		Profile_pic:  artist.Profile_pic,
	}
}
