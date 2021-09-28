package playlists

import (
	"context"
	"myuseek/drivers/databases/songs"
	"time"
)

type Domain struct {
	Id          int
	Name        string
	Description string
	Creator_id  int
	Songs       []songs.Song
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	GetbyID(ctx context.Context, domain Domain) (Domain, error)
	GetPlaylists(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Create(ctx context.Context, domain Domain) (Domain, error)
	GetbyID(ctx context.Context, domain Domain) (Domain, error)
	GetPlaylists(ctx context.Context) ([]Domain, error)
}
