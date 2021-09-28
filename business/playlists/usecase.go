package playlists

import (
	"context"
	"errors"
	"time"
)

type PlaylistUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewPlaylistUsecase(repo Repository, timeout time.Duration) Usecase {
	return &PlaylistUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

//core business register
func (uc *PlaylistUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {

	if domain.Name == "" {
		return Domain{}, errors.New("name empty")
	}
	if domain.Creator_id == 0 {
		return Domain{}, errors.New("creator_id empty")
	}

	playlist, err := uc.Repo.Create(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return playlist, nil
}

func (uc *PlaylistUsecase) GetPlaylists(ctx context.Context) ([]Domain, error) {

	playlistlistdomain, err := uc.Repo.GetPlaylists(ctx)

	if err != nil {
		return []Domain{}, err
	}

	return playlistlistdomain, nil
}
