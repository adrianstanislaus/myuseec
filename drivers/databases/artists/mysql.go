package artists

import (
	"context"
	"myuseek/business/artists"

	"gorm.io/gorm"
)

type MysqlArtistRepository struct {
	Conn *gorm.DB
}

func NewMysqlArtistRepository(conn *gorm.DB) artists.Repository {
	return &MysqlArtistRepository{
		Conn: conn,
	}
}

func (rep *MysqlArtistRepository) Register(ctx context.Context, domain artists.Domain) (artists.Domain, error) {
	artistDB := FromDomain(domain)
	result := rep.Conn.Create(&artistDB)

	if result.Error != nil {
		return artists.Domain{}, result.Error
	}

	return artistDB.ToDomain(), nil
}

func (rep *MysqlArtistRepository) GetArtists(ctx context.Context, domain artists.Domain) ([]artists.Domain, error) {
	artistlist := []Artist{}
	search := "%" + domain.Name + "%"
	result := rep.Conn.Where("name LIKE ?", search).Find(&artistlist)

	if result.Error != nil {
		artistlistdomain := []artists.Domain{}
		return artistlistdomain, result.Error
	}
	artistlistdomain := ToListDomain(artistlist)
	return artistlistdomain, nil

}

func (rep *MysqlArtistRepository) GetArtistById(ctx context.Context, domain artists.Domain) (artists.Domain, error) {
	artist := Artist{}
	result := rep.Conn.Find(&artist, domain.Id)

	if result.Error != nil {
		artistdomain := artists.Domain{}
		return artistdomain, result.Error
	}
	artistdomain := artist.ToDomain()
	return artistdomain, nil

}
