package playlists_test

import (
	"context"
	"myuseek/business/playlists"
	_mockPlaylistRepository "myuseek/business/playlists/mocks"
	"myuseek/business/songs"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var playlistRepository _mockPlaylistRepository.Repository

var playlistService playlists.Usecase
var playlistDomain playlists.Domain
var playlistlistDomain []playlists.Domain

func setup() {
	playlistService = playlists.NewPlaylistUsecase(&playlistRepository, time.Hour*1)
	playlistDomain = playlists.Domain{
		Id:          1,
		Name:        "Judul Playlist",
		Description: "Deskripsi playlist",
		Songs:       []songs.Domain{{Id: 1}},
		Creator_id:  1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	playlistlistDomain = []playlists.Domain{{
		Id:          1,
		Name:        "Judul Playlist",
		Description: "Deskripsi playlist",
		Songs:       []songs.Domain{},
		Creator_id:  1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now()},
	}
}

func TestCreate(t *testing.T) {
	setup()
	playlistRepository.On("Create",
		mock.Anything,
		mock.AnythingOfType("Domain"),
	).Return(playlistDomain, nil).Once()

	t.Run("Test Case 1 | Valid Create", func(t *testing.T) {
		_, err := playlistService.Create(context.Background(), playlists.Domain{
			Name:        "Judul Playlist",
			Creator_id:  1,
			Description: "Deskripsi playlist",
		})
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Invalid Create - Name Empty", func(t *testing.T) {
		_, err := playlistService.Create(context.Background(), playlists.Domain{
			Name:        "",
			Creator_id:  1,
			Description: "Deskripsi playlist",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 3 | Invalid Create - Creator_id Empty", func(t *testing.T) {
		_, err := playlistService.Create(context.Background(), playlists.Domain{
			Name:        "Judul Playlist",
			Creator_id:  0,
			Description: "Deskripsi playlist",
		})
		assert.NotNil(t, err)
	})

}

func TestGetbyID(t *testing.T) {
	setup()
	playlistRepository.On("GetbyID",
		mock.Anything,
		mock.AnythingOfType("Domain"),
	).Return(playlistDomain, nil).Once()

	t.Run("Test Case 1 | Valid GetbyID", func(t *testing.T) {
		_, err := playlistService.GetbyID(context.Background(), playlists.Domain{
			Id: 1,
		})
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Invalid GetbyID - Id Empty", func(t *testing.T) {
		_, err := playlistService.GetbyID(context.Background(), playlists.Domain{
			Id: 0,
		})
		assert.NotNil(t, err)
	})
}

func TestGetPlaylists(t *testing.T) {
	setup()
	playlistRepository.On("GetPlaylists",
		mock.Anything,
	).Return(playlistlistDomain, nil).Once()

	t.Run("Test Case 1 | Valid GetPlaylists", func(t *testing.T) {
		_, err := playlistService.GetPlaylists(context.Background(), playlists.Domain{})
		assert.Nil(t, err)
	})

}

func TestAddSong(t *testing.T) {
	setup()
	playlistRepository.On("AddSong",
		mock.Anything,
		mock.AnythingOfType("Domain"),
	).Return(playlistDomain, nil).Once()

	t.Run("Test Case 1 | Valid AddSong", func(t *testing.T) {
		_, err := playlistService.AddSong(context.Background(), playlists.Domain{
			Id:    1,
			Songs: []songs.Domain{{Id: 1}},
		})
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Invalid AddSong - Id Empty", func(t *testing.T) {
		_, err := playlistService.AddSong(context.Background(), playlists.Domain{
			Id:    0,
			Songs: []songs.Domain{{Id: 1}},
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 2 | Invalid AddSong - Songs Empty", func(t *testing.T) {
		_, err := playlistService.AddSong(context.Background(), playlists.Domain{
			Id:    1,
			Songs: []songs.Domain{},
		})
		assert.NotNil(t, err)
	})
}

func TestRemoveSong(t *testing.T) {
	setup()
	playlistRepository.On("RemoveSong",
		mock.Anything,
		mock.AnythingOfType("Domain"),
	).Return(playlistDomain, nil).Once()

	t.Run("Test Case 1 | Valid RemoveSong", func(t *testing.T) {
		_, err := playlistService.RemoveSong(context.Background(), playlists.Domain{
			Id:    1,
			Songs: []songs.Domain{{Id: 1}},
		})
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Invalid RemoveSong - Id Empty", func(t *testing.T) {
		_, err := playlistService.RemoveSong(context.Background(), playlists.Domain{
			Id:    0,
			Songs: []songs.Domain{{Id: 1}},
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 2 | Invalid RemoveSong - Songs Empty", func(t *testing.T) {
		_, err := playlistService.RemoveSong(context.Background(), playlists.Domain{
			Id:    1,
			Songs: []songs.Domain{},
		})
		assert.NotNil(t, err)
	})
}
