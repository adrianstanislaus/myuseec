package songs

import (
	"context"
	"myuseek/business/songs"
	lyricsapi "myuseek/drivers/databases/thirdparty"

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
	artist := Artist{}
	result := rep.Conn.Create(&songDB)

	if result.Error != nil {
		return songs.Domain{}, result.Error
	}

	result_artist := rep.Conn.Find(&artist, songDB.Artist_id)

	if result_artist.Error != nil {
		songdomain := songs.Domain{}
		return songdomain, result_artist.Error
	}

	songDB.Artist = artist
	return songDB.ToDomain(), nil
}

func (rep *MysqlSongRepository) GetSongs(ctx context.Context, domain songs.Domain) ([]songs.Domain, error) {
	songlist := []Song{}
	titlesearch := "%" + domain.Title + "%"
	result := rep.Conn.Where("Title LIKE ?", titlesearch).Find(&songlist)
	songlistwithArtist := []Song{}

	for _, song := range songlist {
		artist := Artist{}
		result_artist := rep.Conn.Find(&artist, song.Artist_id)

		if result_artist.Error != nil {
			songlistdomain := []songs.Domain{}
			return songlistdomain, result_artist.Error
		}

		song.Artist = artist
		songlistwithArtist = append(songlistwithArtist, song)
	}

	if result.Error != nil {
		songlistdomain := []songs.Domain{}
		return songlistdomain, result.Error
	}
	songlistdomain := ToListDomain(songlistwithArtist)
	return songlistdomain, nil

}

func (rep *MysqlSongRepository) GetSongById(ctx context.Context, domain songs.Domain) (songs.Domain, error) {
	song := FromDomain(domain)
	artist := Artist{}
	result := rep.Conn.Find(&song)

	if result.Error != nil {
		songdomain := songs.Domain{}
		return songdomain, result.Error
	}

	result_artist := rep.Conn.Find(&artist, song.Artist_id)

	if result_artist.Error != nil {
		songdomain := songs.Domain{}
		return songdomain, result_artist.Error
	}
	song.Artist = artist
	songdomain := song.ToDomain()
	return songdomain, nil

}

func (rep *MysqlSongRepository) GetSongLyrics(ctx context.Context, domain songs.Domain) (songs.Domain, error) {
	song := FromDomain(domain)
	artist := Artist{}
	result := rep.Conn.Find(&song)
	result_artist := rep.Conn.Find(&artist, song.Artist_id)

	if result.Error != nil {
		songdomain := songs.Domain{}
		return songdomain, result.Error
	}

	if result_artist.Error != nil {
		songdomain := songs.Domain{}
		return songdomain, result_artist.Error
	}
	song.Artist = artist

	if song.Lyrics == "" {
		lyrics, err_api := lyricsapi.GetLyrics(song.Title, song.Artist.Name)

		if err_api != nil {
			songdomain := songs.Domain{}
			return songdomain, err_api
		}

		rep.Conn.Model(&song).Update("lyrics", lyrics)
	}

	songdomain := song.ToDomain()
	return songdomain, nil

}
