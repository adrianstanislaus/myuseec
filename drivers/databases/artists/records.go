package artists

import (
	"myuseek/business/artists"
	"myuseek/drivers/databases/songs"
	"time"

	"gorm.io/gorm"
)

type Artist struct {
	Id           int    `gorm:"primaryKey"`
	Name         string `gorm:"unique"`
	About        string
	Record_label string
	Profile_pic  string
	Songs        []songs.Song `gorm:"foreignKey:Artist_id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (artist *Artist) ToDomain() artists.Domain {
	return artists.Domain{
		Id:           artist.Id,
		Name:         artist.Name,
		About:        artist.About,
		Record_label: artist.Record_label,
		Profile_pic:  artist.Profile_pic,
		CreatedAt:    artist.CreatedAt,
		UpdatedAt:    artist.UpdatedAt,
	}
}

func FromDomain(domain artists.Domain) Artist {
	return Artist{
		Id:           domain.Id,
		Name:         domain.Name,
		About:        domain.About,
		Record_label: domain.Record_label,
		Profile_pic:  domain.Profile_pic,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt}
}

func ToListDomain(data []Artist) (result []artists.Domain) {
	result = []artists.Domain{}
	for _, artist := range data {
		result = append(result, artist.ToDomain())
	}
	return result
}
