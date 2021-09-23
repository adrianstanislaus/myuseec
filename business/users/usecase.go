package users

import (
	"context"
	"errors"
	"time"
)

type UserUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repo Repository, timeout time.Duration) Usecase {
	return &UserUsecase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

//core business register
func (uc *UserUsecase) Register(ctx context.Context, firstname, lastname, username, email, password, bio, profile_pic, subscription_type string) (Domain, error) {

	if firstname == "" {
		return Domain{}, errors.New("email empty")
	}
	if lastname == "" {
		return Domain{}, errors.New("email empty")
	}

	if username == "" {
		return Domain{}, errors.New("email empty")
	}

	if email == "" {
		return Domain{}, errors.New("email empty")
	}

	if password == "" {
		return Domain{}, errors.New("email empty")
	}

	if subscription_type == "" {
		return Domain{}, errors.New("email empty")
	}

	user, err := uc.Repo.Register(ctx, firstname, lastname, username, email, password, bio, profile_pic, subscription_type)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
