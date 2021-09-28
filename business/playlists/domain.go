package playlists

import (
	"context"
	"myuseek/business/songs"
	"time"
)

type Domain struct {
	Id          int
	Name        string
	Description string
	Creator_id  int
	Songs       []songs.Domain
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	GetbyID(ctx context.Context, domain Domain) (Domain, error)
	AddSong(ctx context.Context, domain Domain) (Domain, error)
	GetPlaylists(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	GetbyID(ctx context.Context, domain Domain) (Domain, error)
	AddSong(ctx context.Context, domain Domain) (Domain, error)
	GetPlaylists(ctx context.Context) ([]Domain, error)
}
