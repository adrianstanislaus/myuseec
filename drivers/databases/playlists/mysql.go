package playlists

import (
	"context"
	"myuseek/business/playlists"

	"gorm.io/gorm"
)

type MysqlPlaylistRepository struct {
	Conn *gorm.DB
}

func NewMysqlPlaylistRepository(conn *gorm.DB) playlists.Repository {
	return &MysqlPlaylistRepository{
		Conn: conn,
	}
}

func (rep *MysqlPlaylistRepository) Create(ctx context.Context, domain playlists.Domain) (playlists.Domain, error) {
	playlistDB := FromDomain(domain)
	result := rep.Conn.Create(&playlistDB)

	if result.Error != nil {
		return playlists.Domain{}, result.Error
	}

	return playlistDB.ToDomain(), nil
}

func (rep *MysqlPlaylistRepository) GetbyID(ctx context.Context, domain playlists.Domain) (playlists.Domain, error) {
	playlist := Playlist{}
	result := rep.Conn.Find(&playlist, domain.Id)

	if result.Error != nil {
		playlistdomain := playlists.Domain{}
		return playlistdomain, result.Error
	}
	playlistdomain := playlist.ToDomain()
	return playlistdomain, nil

}
func (rep *MysqlPlaylistRepository) GetPlaylists(ctx context.Context) ([]playlists.Domain, error) {
	playlistlist := []Playlist{}
	result := rep.Conn.Find(&playlistlist)

	if result.Error != nil {
		playlistlistdomain := []playlists.Domain{}
		return playlistlistdomain, result.Error
	}
	playlistlistdomain := ToListDomain(playlistlist)
	return playlistlistdomain, nil

}
