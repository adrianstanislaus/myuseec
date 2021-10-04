package albums_test

import (
	"context"
	"myuseek/business/albums"
	_mockAlbumRepository "myuseek/business/albums/mocks"
	"myuseek/business/songs"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var albumRepository _mockAlbumRepository.Repository

var albumService albums.Usecase
var albumDomain albums.Domain

func setup() {
	albumService = albums.NewAlbumUsecase(&albumRepository, time.Hour*1)
	albumDomain = albums.Domain{
		Id:           1,
		Title:        "Nama Album",
		Artist_id:    1,
		Album_art:    "linkkegambaralbum.com",
		Songs:        []songs.Domain{},
		Release_date: time.Now(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func TestRegister(t *testing.T) {
	setup()
	albumRepository.On("Add",
		mock.Anything,
		mock.AnythingOfType("Domain"),
	).Return(albumDomain, nil).Once()

	t.Run("Test Case 1 | Valid Add", func(t *testing.T) {
		_, err := albumService.Add(context.Background(), albums.Domain{
			Title:        "Judul Album",
			Artist_id:    1,
			Album_art:    "linkalbumart.com",
			Songs:        []songs.Domain{},
			Release_date: time.Now(),
		})
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Invalid Add - Title Empty ", func(t *testing.T) {
		_, err := albumService.Add(context.Background(), albums.Domain{
			Title:        "",
			Artist_id:    1,
			Album_art:    "linkalbumart.com",
			Songs:        []songs.Domain{},
			Release_date: time.Now(),
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 3 | Invalid Add - Artist_id Empty ", func(t *testing.T) {
		_, err := albumService.Add(context.Background(), albums.Domain{
			Title:        "Judul Album",
			Artist_id:    0,
			Album_art:    "linkalbumart.com",
			Songs:        []songs.Domain{},
			Release_date: time.Now(),
		})
		assert.NotNil(t, err)
	})
}
