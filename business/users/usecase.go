package users

import (
	"context"
	"errors"
	"myuseek/app/middlewares"
	encrypt "myuseek/helpers/hashing"
	"time"
)

type UserUsecase struct {
	ConfigJWT      middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUsecase(configJWT middlewares.ConfigJWT, repo Repository, timeout time.Duration) Usecase {
	return &UserUsecase{
		ConfigJWT:      configJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

//core business register
func (uc *UserUsecase) Register(ctx context.Context, domain Domain) (Domain, error) {

	if domain.FirstName == "" {
		return Domain{}, errors.New("firstname empty")
	}
	if domain.LastName == "" {
		return Domain{}, errors.New("lastname empty")
	}

	if domain.Username == "" {
		return Domain{}, errors.New("username empty")
	}

	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}

	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}

	if domain.Subscription_type == "" {
		return Domain{}, errors.New("subscription type empty")
	}

	//encryption
	var err error
	domain.Password, err = encrypt.Hash(domain.Password)

	user, err := uc.Repo.Register(ctx, domain.FirstName, domain.LastName, domain.Username, domain.Email, domain.Password, domain.Bio, domain.Profile_pic, domain.Subscription_type)

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

	user, err := uc.Repo.Login(ctx, domain.Username, domain.Password)

	if err != nil {
		return Domain{}, err
	}

	if !(encrypt.CheckPasswordHash(domain.Password, user.Password)) {
		return Domain{}, errors.New("wrong password")
	}

	user.Token, err = uc.ConfigJWT.GenerateToken(user.Id)
	return user, nil
}

func (uc *UserUsecase) GetUsers(ctx context.Context) ([]Domain, error) {

	userlistdomain, err := uc.Repo.GetUsers(ctx)

	if err != nil {
		return []Domain{}, err
	}

	return userlistdomain, nil
}
