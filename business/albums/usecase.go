package albums

import (
	"context"
	"errors"
	"time"
)

type AlbumUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewAlbumUsecase(repo Repository, timeout time.Duration) Usecase {
	return &AlbumUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

//core business register
func (uc *AlbumUsecase) Add(ctx context.Context, domain Domain) (Domain, error) {

	if domain.Title == "" {
		return Domain{}, errors.New("title empty")
	}
	if domain.Artist_id == 0 {
		return Domain{}, errors.New("artist_id empty")
	}

	album, err := uc.Repo.Add(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return album, nil
}

func (uc *AlbumUsecase) GetAlbums(ctx context.Context, domain Domain) ([]Domain, error) {

	albumlistdomain, err := uc.Repo.GetAlbums(ctx, domain)

	if err != nil {
		return []Domain{}, err
	}

	return albumlistdomain, nil
}

func (uc *AlbumUsecase) GetAlbumById(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Id == 0 {
		return Domain{}, errors.New("id empty")
	}

	albumdomain, err := uc.Repo.GetAlbumById(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	if albumdomain.Title == "" {
		return Domain{}, errors.New("no album found with that ID")
	}

	return albumdomain, nil
}
