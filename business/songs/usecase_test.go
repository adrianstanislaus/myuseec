package songs_test

import (
	"context"
	"myuseek/business/artists"
	"myuseek/business/songs"
	_mockSongRepository "myuseek/business/songs/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var songRepository _mockSongRepository.Repository

var songService songs.Usecase
var songDomain songs.Domain
var songlistDomain []songs.Domain

func setup() {
	songService = songs.NewSongUsecase(&songRepository, time.Hour*1)
	songDomain = songs.Domain{
		Id:        1,
		Title:     "Judul Lagu",
		Artist_id: 1,
		Artist:    artists.Domain{},
		Album_id:  1,
		Genre:     "Jazz",
		Duration:  "human being",
		Lyrics:    "lalala lilili",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func TestAdd(t *testing.T) {
	setup()
	songRepository.On("Add",
		mock.Anything,
		mock.AnythingOfType("Domain"),
	).Return(songDomain, nil).Once()

	t.Run("Test Case 1 | Valid Add", func(t *testing.T) {
		_, err := songService.Add(context.Background(), songs.Domain{
			Title:     "Judul Lagu",
			Artist_id: 1,
			Artist:    artists.Domain{},
			Album_id:  1,
			Genre:     "Jazz",
			Duration:  "human being",
			Lyrics:    "lalala lilili",
		})
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Invalid Add - Title Empty", func(t *testing.T) {
		_, err := songService.Add(context.Background(), songs.Domain{
			Title:     "",
			Artist_id: 1,
			Artist:    artists.Domain{},
			Album_id:  1,
			Genre:     "Jazz",
			Duration:  "human being",
			Lyrics:    "lalala lilili",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 3 | Invalid Add - Artist_id Empty", func(t *testing.T) {
		_, err := songService.Add(context.Background(), songs.Domain{
			Title:     "Judul Lagu",
			Artist_id: 0,
			Artist:    artists.Domain{},
			Album_id:  1,
			Genre:     "Jazz",
			Duration:  "human being",
			Lyrics:    "lalala lilili",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 4 | Invalid Add - Album_id Empty", func(t *testing.T) {
		_, err := songService.Add(context.Background(), songs.Domain{
			Title:     "Judul lagu",
			Artist_id: 1,
			Artist:    artists.Domain{},
			Album_id:  0,
			Genre:     "Jazz",
			Duration:  "human being",
			Lyrics:    "lalala lilili",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 5 | Invalid Add - Duration Empty", func(t *testing.T) {
		_, err := songService.Add(context.Background(), songs.Domain{
			Title:     "Judul Lagu",
			Artist_id: 1,
			Artist:    artists.Domain{},
			Album_id:  1,
			Genre:     "Jazz",
			Duration:  "",
			Lyrics:    "lalala lilili",
		})
		assert.NotNil(t, err)
	})

}

func TestGetSongById(t *testing.T) {
	setup()
	songRepository.On("GetSongById",
		mock.Anything,
		mock.AnythingOfType("Domain"),
	).Return(songDomain, nil).Once()

	t.Run("Test Case 1 | Valid GetSongById", func(t *testing.T) {
		_, err := songService.GetSongById(context.Background(), songs.Domain{
			Id: 1,
		})
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Invalid GetSongById - No Id ", func(t *testing.T) {
		_, err := songService.GetSongById(context.Background(), songs.Domain{
			Id: 0,
		})
		assert.NotNil(t, err)
	})

}

func TestGetSongs(t *testing.T) {
	setup()
	songRepository.On("GetSongs",
		mock.Anything,
	).Return(songlistDomain, nil).Once()

	t.Run("Test Case 1 | Valid GetSongs", func(t *testing.T) {
		_, err := songService.GetSongs(context.Background())
		assert.Nil(t, err)
	})

}

func TestGetSongLyrics(t *testing.T) {
	setup()
	songRepository.On("GetSongLyrics",
		mock.Anything,
		mock.AnythingOfType("Domain"),
	).Return(songDomain, nil).Once()

	t.Run("Test Case 1 | Valid GetSongLyrics", func(t *testing.T) {
		_, err := songService.GetSongLyrics(context.Background(), songs.Domain{
			Id: 1,
		})
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Invalid GetSongLyrics - No Id ", func(t *testing.T) {
		_, err := songService.GetSongLyrics(context.Background(), songs.Domain{
			Id: 0,
		})
		assert.NotNil(t, err)
	})

}
