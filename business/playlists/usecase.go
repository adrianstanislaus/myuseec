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

func (uc *PlaylistUsecase) GetbyID(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Id == 0 {
		return Domain{}, errors.New("id empty")
	}

	playlistdomain, err := uc.Repo.GetbyID(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	if playlistdomain.Name == "" {
		return Domain{}, errors.New("no playlist found with that ID")
	}

	return playlistdomain, nil
}

func (uc *PlaylistUsecase) AddSong(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Id == 0 {
		return Domain{}, errors.New("id empty")
	}
	if len(domain.Songs) == 0 {
		return Domain{}, errors.New("song_to_add empty")
	}
	playlistdomain, err := uc.Repo.AddSong(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	if playlistdomain.Name == "" {
		return Domain{}, errors.New("no playlist found with that ID")
	}

	if len(playlistdomain.Songs) == 0 {
		return Domain{}, errors.New("song not found in playlist")
	}

	return playlistdomain, nil
}

func (uc *PlaylistUsecase) RemoveSong(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Id == 0 {
		return Domain{}, errors.New("id empty")
	}
	if len(domain.Songs) == 0 {
		return Domain{}, errors.New("song_to_remove empty")
	}
	playlistdomain, err := uc.Repo.RemoveSong(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	if playlistdomain.Name == "" {
		return Domain{}, errors.New("no playlist found with that ID")
	}

	if len(playlistdomain.Songs) == 0 {
		return Domain{}, errors.New("song not found in playlist")
	}

	return playlistdomain, nil
}

func (uc *PlaylistUsecase) GetPlaylists(ctx context.Context, domain Domain) ([]Domain, error) {

	playlistlistdomain, err := uc.Repo.GetPlaylists(ctx, domain)

	if err != nil {
		return []Domain{}, err
	}

	return playlistlistdomain, nil
}
