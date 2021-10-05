package artists_test

import (
	"context"
	"myuseek/business/artists"
	_mockArtistRepository "myuseek/business/artists/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var artistRepository _mockArtistRepository.Repository

var artistService artists.Usecase
var artistDomain artists.Domain
var artistlistDomain []artists.Domain

func setup() {
	artistService = artists.NewArtistUsecase(&artistRepository, time.Hour*1)
	artistDomain = artists.Domain{
		Id:           1,
		Name:         "Nama Artist",
		About:        "Tentang si Artist",
		Record_label: "record label si artist",
		Profile_pic:  "linkprofilefoto.com",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	artistlistDomain = []artists.Domain{{
		Id:           1,
		Name:         "Nama Artist",
		About:        "Tentang si Artist",
		Record_label: "record label si artist",
		Profile_pic:  "linkprofilefoto.com",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now()},
	}
}

func TestRegister(t *testing.T) {
	setup()
	artistRepository.On("Register",
		mock.Anything,
		mock.AnythingOfType("Domain"),
	).Return(artistDomain, nil).Once()

	t.Run("Test Case 1 | Valid Register", func(t *testing.T) {
		_, err := artistService.Register(context.Background(), artists.Domain{
			Name:         "Nama Artist",
			About:        "Tentang artist",
			Record_label: "record label si artist",
			Profile_pic:  "linkkefoto.com",
		})
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Invalid Register - Name Empty", func(t *testing.T) {
		_, err := artistService.Register(context.Background(), artists.Domain{
			Name:         "",
			About:        "Tentang artist",
			Record_label: "record label si artist",
			Profile_pic:  "linkkefoto.com",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 3 | Invalid Create - About Empty", func(t *testing.T) {
		_, err := artistService.Register(context.Background(), artists.Domain{
			Name:         "Nama Artist",
			About:        "",
			Record_label: "record label si artist",
			Profile_pic:  "linkkefoto.com",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 4 | Invalid Create - Record_label Empty", func(t *testing.T) {
		_, err := artistService.Register(context.Background(), artists.Domain{
			Name:         "Nama Artist",
			About:        "Tentang artist",
			Record_label: "",
			Profile_pic:  "linkkefoto.com",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 5 | Invalid Create - Profile_pic Empty", func(t *testing.T) {
		_, err := artistService.Register(context.Background(), artists.Domain{
			Name:         "Nama Artist",
			About:        "Tentang artist",
			Record_label: "record label si artist",
			Profile_pic:  "",
		})
		assert.NotNil(t, err)
	})

}

func TestGetArtists(t *testing.T) {
	setup()
	artistRepository.On("GetArtists",
		mock.Anything,
	).Return(artistlistDomain, nil).Once()

	t.Run("Test Case 1 | Valid GetArtists", func(t *testing.T) {
		_, err := artistService.GetArtists(context.Background())
		assert.Nil(t, err)
	})

}

func TestGetArtistById(t *testing.T) {
	setup()
	artistRepository.On("GetArtistById",
		mock.Anything,
		mock.AnythingOfType("Domain"),
	).Return(artistDomain, nil).Once()

	t.Run("Test Case 1 | Valid GetArtistById", func(t *testing.T) {
		_, err := artistService.GetArtistById(context.Background(), artists.Domain{
			Id: 1,
		})
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Invalid GetAlbumById - Id Empty", func(t *testing.T) {
		_, err := artistService.GetArtistById(context.Background(), artists.Domain{
			Id: 0,
		})
		assert.NotNil(t, err)
	})
}
