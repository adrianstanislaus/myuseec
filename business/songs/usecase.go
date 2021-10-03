package songs

import (
	"context"
	"errors"
	"time"
)

type SongUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewSongUsecase(repo Repository, timeout time.Duration) Usecase {
	return &SongUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

//core business register
func (uc *SongUsecase) Add(ctx context.Context, domain Domain) (Domain, error) {

	if domain.Title == "" {
		return Domain{}, errors.New("title empty")
	}
	if domain.Artist_id == 0 {
		return Domain{}, errors.New("artist_id empty")
	}

	if domain.Album_id == 0 {
		return Domain{}, errors.New("album_id empty")
	}

	if domain.Duration == "" {
		return Domain{}, errors.New("duration empty")
	}

	song, err := uc.Repo.Add(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return song, nil
}

func (uc *SongUsecase) GetSongs(ctx context.Context) ([]Domain, error) {

	songlistdomain, err := uc.Repo.GetSongs(ctx)

	if err != nil {
		return []Domain{}, err
	}

	return songlistdomain, nil
}

func (uc *SongUsecase) GetSongById(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Id == 0 {
		return Domain{}, errors.New("id empty")
	}

	songdomain, err := uc.Repo.GetSongById(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	if songdomain.Title == "" {
		return Domain{}, errors.New("no song found with that ID")
	}

	return songdomain, nil
}

func (uc *SongUsecase) GetSongLyrics(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Id == 0 {
		return Domain{}, errors.New("id empty")
	}

	songdomain, err := uc.Repo.GetSongLyrics(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	if songdomain.Title == "" {
		return Domain{}, errors.New("no song found with that ID")
	}

	return songdomain, nil
}
