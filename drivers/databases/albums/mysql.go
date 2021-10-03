package albums

import (
	"context"
	"myuseek/business/albums"

	"gorm.io/gorm"
)

type MysqlAlbumRepository struct {
	Conn *gorm.DB
}

func NewMysqlAlbumRepository(conn *gorm.DB) albums.Repository {
	return &MysqlAlbumRepository{
		Conn: conn,
	}
}

func (rep *MysqlAlbumRepository) Add(ctx context.Context, domain albums.Domain) (albums.Domain, error) {
	albumDB := FromDomain(domain)
	result := rep.Conn.Create(&albumDB)

	if result.Error != nil {
		return albums.Domain{}, result.Error
	}

	return albumDB.ToDomain(), nil
}

func (rep *MysqlAlbumRepository) GetAlbumById(ctx context.Context, domain albums.Domain) (albums.Domain, error) {
	album := Album{}
	result := rep.Conn.Preload("Songs").Find(&album, domain.Id)

	if result.Error != nil {
		albumdomain := albums.Domain{}
		return albumdomain, result.Error
	}
	albumdomain := album.ToDomain()
	return albumdomain, nil

}

func (rep *MysqlAlbumRepository) GetAlbums(ctx context.Context) ([]albums.Domain, error) {
	albumlist := []Album{}
	result := rep.Conn.Find(&albumlist)

	if result.Error != nil {
		albumlistdomain := []albums.Domain{}
		return albumlistdomain, result.Error
	}
	albumlistdomain := ToListDomain(albumlist)
	return albumlistdomain, nil

}
