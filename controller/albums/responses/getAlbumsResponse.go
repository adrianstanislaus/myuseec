package responses

import (
	"myuseek/business/albums"
)

type GetAlbumsResponse struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Artist_id int    `json:"artist_id"`
}

func FromDomainToGetAlbums(domain albums.Domain) GetAlbumsResponse {
	return GetAlbumsResponse{
		Id:        domain.Id,
		Title:     domain.Title,
		Artist_id: domain.Artist_id,
	}
}

func FromListDomainToGetAlbums(data []albums.Domain) (result []GetAlbumsResponse) {
	result = []GetAlbumsResponse{}
	for _, album := range data {
		result = append(result, FromDomainToGetAlbums(album))
	}
	return result
}
