package songs

import (
	"context"
	"time"
)

type Domain struct {
	Id        int
	Title     string
	Artist_id int
	Album_id  int
	Genre     string
	Duration  string
	Lyrics    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetSongById(ctx context.Context, domain Domain) (Domain, error)
	GetSongs(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Add(ctx context.Context, domain Domain) (Domain, error)
	GetSongById(ctx context.Context, domain Domain) (Domain, error)
	GetSongs(ctx context.Context) ([]Domain, error)
}
