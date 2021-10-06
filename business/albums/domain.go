package albums

import (
	"context"
	"myuseek/business/songs"
	"time"
)

type Domain struct {
	Id           int
	Title        string
	Artist_id    int
	Songs        []songs.Domain
	Album_art    string
	Release_date time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Usecase interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAlbumById(ctx context.Context, domain Domain) (Domain, error)
	GetAlbums(ctx context.Context, domain Domain) ([]Domain, error)
}

type Repository interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetAlbumById(ctx context.Context, domain Domain) (Domain, error)
	GetAlbums(ctx context.Context, domain Domain) ([]Domain, error)
}
