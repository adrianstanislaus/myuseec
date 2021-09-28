package responses

import (
	"myuseek/business/artists"
	"time"
)

type ArtistResponse struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	About        string    `json:"about"`
	Record_label string    `json:"record_label"`
	Profile_pic  string    `json:"profile_pic"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func FromDomain(domain artists.Domain) ArtistResponse {
	return ArtistResponse{
		Id:           domain.Id,
		Name:         domain.Name,
		About:        domain.About,
		Record_label: domain.Record_label,
		Profile_pic:  domain.Profile_pic,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}

func FromListDomain(data []artists.Domain) (result []ArtistResponse) {
	result = []ArtistResponse{}
	for _, artist := range data {
		result = append(result, FromDomain(artist))
	}
	return result
}
