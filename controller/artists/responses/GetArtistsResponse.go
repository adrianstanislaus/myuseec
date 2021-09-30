package responses

import (
	"myuseek/business/artists"
)

type GetArtistsResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Profile_pic string `json:"profile_pic"`
}

func FromDomainToGetArtists(domain artists.Domain) GetArtistsResponse {
	return GetArtistsResponse{
		Id:          domain.Id,
		Name:        domain.Name,
		Profile_pic: domain.Profile_pic,
	}
}

func FromListDomainToGetArtists(data []artists.Domain) (result []GetArtistsResponse) {
	result = []GetArtistsResponse{}
	for _, artist := range data {
		result = append(result, FromDomainToGetArtists(artist))
	}
	return result
}
