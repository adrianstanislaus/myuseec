package responses

import (
	"myuseek/business/artists"
)

type GetArtistByIdResponse struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	About        string `json:"about"`
	Record_label string `json:"record_label"`
	Profile_pic  string `json:"profile_pic"`
}

func FromDomainToGetArtistById(domain artists.Domain) GetArtistByIdResponse {
	return GetArtistByIdResponse{
		Id:           domain.Id,
		Name:         domain.Name,
		About:        domain.About,
		Record_label: domain.Record_label,
		Profile_pic:  domain.Profile_pic,
	}
}
