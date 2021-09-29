package songs

import (
	"context"
	"myuseek/business/songs"

	"gorm.io/gorm"
)

type MysqlSongRepository struct {
	Conn *gorm.DB
}

func NewMysqlSongRepository(conn *gorm.DB) songs.Repository {
	return &MysqlSongRepository{
		Conn: conn,
	}
}

func (rep *MysqlSongRepository) Add(ctx context.Context, domain songs.Domain) (songs.Domain, error) {
	songDB := FromDomain(domain)
	result := rep.Conn.Create(&songDB)

	if result.Error != nil {
		return songs.Domain{}, result.Error
	}

	return songDB.ToDomain(), nil
}

func (rep *MysqlSongRepository) GetSongs(ctx context.Context) ([]songs.Domain, error) {
	songlist := []Song{}
	result := rep.Conn.Find(&songlist)

	if result.Error != nil {
		songlistdomain := []songs.Domain{}
		return songlistdomain, result.Error
	}
	songlistdomain := ToListDomain(songlist)
	return songlistdomain, nil

}

func (rep *MysqlSongRepository) GetSongById(ctx context.Context, domain songs.Domain) (songs.Domain, error) {
	song := Song{}
	result := rep.Conn.Find(&song, domain.Id)

	if result.Error != nil {
		songdomain := songs.Domain{}
		return songdomain, result.Error
	}
	songdomain := song.ToDomain()
	return songdomain, nil

}
