package artists

import (
	"context"
	"time"
)

type Domain struct {
	Id           int
	Name         string
	About        string
	Record_label string
	Profile_pic  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Usecase interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	GetArtists(ctx context.Context, domain Domain) ([]Domain, error)
	GetArtistById(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	GetArtists(ctx context.Context, domain Domain) ([]Domain, error)
	GetArtistById(ctx context.Context, domain Domain) (Domain, error)
}
