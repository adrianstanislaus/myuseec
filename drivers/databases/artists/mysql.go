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

func (rep *MysqlArtistRepository) GetArtists(ctx context.Context) ([]artists.Domain, error) {
	artistlist := []Artist{}
	result := rep.Conn.Find(&artistlist)

	if result.Error != nil {
		artistlistdomain := []artists.Domain{}
		return artistlistdomain, result.Error
	}
	artistlistdomain := ToListDomain(artistlist)
	return artistlistdomain, nil

}
