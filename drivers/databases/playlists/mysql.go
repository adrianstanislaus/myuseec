package playlists

import (
	"context"
	"myuseek/business/playlists"
	"myuseek/drivers/databases/songs"

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
	result := rep.Conn.Preload("Songs").Find(&playlist, domain.Id)

	if result.Error != nil {
		playlistdomain := playlists.Domain{}
		return playlistdomain, result.Error
	}
	playlistdomain := playlist.ToDomain()
	return playlistdomain, nil

}

func (rep *MysqlPlaylistRepository) AddSong(ctx context.Context, domain playlists.Domain) (playlists.Domain, error) {
	playlist := Playlist{}
	playlist.Id = domain.Id
	song := songs.FromListDomain(domain.Songs)
	errorDb := rep.Conn.Model(&playlist).Association("Songs").Append(song)
	errorDbsong := rep.Conn.Model(&playlist).Where("id = ?", song[0].Id).Association("Songs").Find(&song)

	if errorDb != nil {
		playlistdomain := playlists.Domain{}
		return playlistdomain, errorDb
	}

	if errorDbsong != nil {
		playlistdomain := playlists.Domain{}
		return playlistdomain, errorDbsong
	}

	result := rep.Conn.Find(&playlist, domain.Id)

	if result.Error != nil {
		playlistdomain := playlists.Domain{}
		return playlistdomain, result.Error
	}
	playlist.Songs = song
	playlistdomain := playlist.ToDomain()
	return playlistdomain, nil

}

func (rep *MysqlPlaylistRepository) RemoveSong(ctx context.Context, domain playlists.Domain) (playlists.Domain, error) {
	playlist := Playlist{}
	playlist.Id = domain.Id
	song := songs.FromListDomain(domain.Songs)
	errorDbsong := rep.Conn.Model(&playlist).Where("id = ?", song[0].Id).Association("Songs").Find(&song)
	errorDb := rep.Conn.Model(&playlist).Association("Songs").Delete(song)

	if errorDb != nil {
		playlistdomain := playlists.Domain{}
		return playlistdomain, errorDb
	}

	if errorDbsong != nil {
		playlistdomain := playlists.Domain{}
		return playlistdomain, errorDbsong
	}

	result := rep.Conn.Find(&playlist, domain.Id)

	if result.Error != nil {
		playlistdomain := playlists.Domain{}
		return playlistdomain, result.Error
	}
	playlist.Songs = song
	playlistdomain := playlist.ToDomain()
	return playlistdomain, nil

}

func (rep *MysqlPlaylistRepository) GetPlaylists(ctx context.Context, domain playlists.Domain) ([]playlists.Domain, error) {
	playlistlist := []Playlist{}
	search := "%" + domain.Name + "%"
	result := rep.Conn.Where("name LIKE ?", search).Find(&playlistlist)

	if result.Error != nil {
		playlistlistdomain := []playlists.Domain{}
		return playlistlistdomain, result.Error
	}
	playlistlistdomain := ToListDomain(playlistlist)
	return playlistlistdomain, nil

}
