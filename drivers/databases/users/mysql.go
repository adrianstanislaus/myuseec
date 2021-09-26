package users

import (
	"context"
	"myuseek/business/users"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) Register(ctx context.Context, firstname, lastname, username, email, password, bio, profile_pic, subscription_type string) (users.Domain, error) {
	var userDB Users
	userDB.FirstName = firstname
	userDB.LastName = lastname
	userDB.Email = email
	userDB.Password = password
	userDB.Username = username
	userDB.Bio = bio
	userDB.Profile_pic = profile_pic
	userDB.Subscription_type = subscription_type
	result := rep.Conn.Create(&userDB)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return userDB.ToDomain(), nil
}

func (rep *MysqlUserRepository) Login(ctx context.Context, username, password string) (users.Domain, error) {
	var user Users
	result := rep.Conn.First(&user, "username = ?",
		username)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil

}
