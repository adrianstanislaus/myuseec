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
		return Domain{}, errors.New("firstname empty")
	}
	if lastname == "" {
		return Domain{}, errors.New("lastname empty")
	}

	if username == "" {
		return Domain{}, errors.New("username empty")
	}

	if email == "" {
		return Domain{}, errors.New("email empty")
	}

	if password == "" {
		return Domain{}, errors.New("password empty")
	}

	if subscription_type == "" {
		return Domain{}, errors.New("subscription type empty")
	}

	user, err := uc.Repo.Register(ctx, firstname, lastname, username, email, password, bio, profile_pic, subscription_type)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUsecase) Login(ctx context.Context, domain Domain) (Domain, error) {

	if domain.Username == "" {
		return Domain{}, errors.New("email empty")
	}

	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}
	// var err error
	// domain.Password, err = encrypt.Hash(domain.Password)

	// if err != nil {
	// 	return Domain{}, err
	// }

	user, err := uc.Repo.Login(ctx, domain.Username, domain.Password)

	if err != nil {
		return Domain{}, err
	}

	// user.Token, err = uc.ConfigJWT.GenerateToken(user.Id)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
