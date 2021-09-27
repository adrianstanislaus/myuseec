package artists

import (
	"context"
	"errors"
	"time"
)

type ArtistUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewArtistUsecase(repo Repository, timeout time.Duration) Usecase {
	return &ArtistUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

//core business register
func (uc *ArtistUsecase) Register(ctx context.Context, domain Domain) (Domain, error) {

	if domain.Name == "" {
		return Domain{}, errors.New("name empty")
	}
	if domain.About == "" {
		return Domain{}, errors.New("about empty")
	}

	if domain.Record_label == "" {
		return Domain{}, errors.New("record empty")
	}

	if domain.Profile_pic == "" {
		return Domain{}, errors.New("email empty")
	}

	artist, err := uc.Repo.Register(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return artist, nil
}

func (uc *ArtistUsecase) GetArtists(ctx context.Context) ([]Domain, error) {

	artistlistdomain, err := uc.Repo.GetArtists(ctx)

	if err != nil {
		return []Domain{}, err
	}

	return artistlistdomain, nil
}
