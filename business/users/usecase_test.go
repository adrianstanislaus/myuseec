package users_test

import (
	"context"
	_middlewares "myuseek/app/middlewares"
	"myuseek/business/users"
	_mockUserRepository "myuseek/business/users/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository _mockUserRepository.Repository

var userService users.Usecase
var userDomain users.Domain
var userlistDomain []users.Domain

var configJWT = _middlewares.ConfigJWT{
	SecretJWT:       "123",
	ExpiresDuration: 72,
}

func setup() {
	userService = users.NewUserUsecase(configJWT, &userRepository, time.Hour*1)
	userDomain = users.Domain{
		Id:                1,
		FirstName:         "Adrian",
		LastName:          "Stanislaus",
		Username:          "adrian_sts",
		Email:             "adrian@gmail.com",
		Password:          "$2a$12$vAr7enV44Uu/8R8fW6VQdeC3r/4UjpagXsj2r0bCatAjNUNpPUnkW",
		Bio:               "human being",
		Profile_pic:       "link.com",
		Subscription_type: "Premium",
		Token:             "123",
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
	userlistDomain = []users.Domain{{
		Id:                1,
		FirstName:         "Adrian",
		LastName:          "Stanislaus",
		Username:          "adrian_sts",
		Email:             "adrian@gmail.com",
		Password:          "$2a$12$vAr7enV44Uu/8R8fW6VQdeC3r/4UjpagXsj2r0bCatAjNUNpPUnkW",
		Bio:               "human being",
		Profile_pic:       "link.com",
		Subscription_type: "Premium",
		Token:             "123",
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now()},
	}
}

func TestRegister(t *testing.T) {
	setup()
	userRepository.On("Register",
		mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
	).Return(userDomain, nil).Once()

	t.Run("Test Case 1 | Valid Register", func(t *testing.T) {
		_, err := userService.Register(context.Background(), users.Domain{
			FirstName:         "Adrian",
			LastName:          "Stanislaus",
			Username:          "adrian_sts",
			Email:             "adrian@gmail.com",
			Password:          "abc123",
			Subscription_type: "Premium",
		})
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Invalid Register - FirstName Empty", func(t *testing.T) {
		_, err := userService.Register(context.Background(), users.Domain{
			FirstName:         "",
			LastName:          "Stanislaus",
			Username:          "adrian_sts",
			Email:             "adrian@gmail.com",
			Password:          "abc123",
			Subscription_type: "Premium",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 3 | Invalid Register - LastName Empty", func(t *testing.T) {
		_, err := userService.Register(context.Background(), users.Domain{
			FirstName:         "Adrian",
			LastName:          "",
			Username:          "adrian_sts",
			Email:             "adrian@gmail.com",
			Password:          "abc123",
			Subscription_type: "Premium",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 4 | Invalid Register - Username Empty", func(t *testing.T) {
		_, err := userService.Register(context.Background(), users.Domain{
			FirstName:         "Adrian",
			LastName:          "Stanislaus",
			Username:          "",
			Email:             "adrian@gmail.com",
			Password:          "abc123",
			Subscription_type: "Premium",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 5 | Invalid Register - Email Empty", func(t *testing.T) {
		_, err := userService.Register(context.Background(), users.Domain{
			FirstName:         "Adrian",
			LastName:          "Stanislaus",
			Username:          "adrian_sts",
			Email:             "",
			Password:          "abc123",
			Subscription_type: "Premium",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 6 | Invalid Register - Password Empty", func(t *testing.T) {
		_, err := userService.Register(context.Background(), users.Domain{
			FirstName:         "Adrian",
			LastName:          "Stanislaus",
			Username:          "adrian_sts",
			Email:             "adrian@gmail.com",
			Password:          "",
			Subscription_type: "Premium",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 7 | Invalid Register - Subscription_type Empty", func(t *testing.T) {
		_, err := userService.Register(context.Background(), users.Domain{
			FirstName:         "Adrian",
			LastName:          "Stanislaus",
			Username:          "adrian_sts",
			Email:             "adrian@gmail.com",
			Password:          "123",
			Subscription_type: "",
		})
		assert.NotNil(t, err)
	})

}

func TestLogin(t *testing.T) {
	setup()
	userRepository.On("Login",
		mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string"),
	).Return(userDomain, nil).Once()

	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		_, err := userService.Login(context.Background(), users.Domain{
			Username: "adrian_sts",
			Password: "abc123",
		})
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Invalid Login - Username Empty", func(t *testing.T) {
		_, err := userService.Login(context.Background(), users.Domain{
			Username: "",
			Password: "abc123",
		})
		assert.NotNil(t, err)
	})

	t.Run("Test Case 3 | Invalid Login - Password Empty", func(t *testing.T) {
		_, err := userService.Login(context.Background(), users.Domain{
			Username: "adrian_sts",
			Password: "",
		})
		assert.NotNil(t, err)
	})

}

func TestGetUsers(t *testing.T) {
	setup()
	userRepository.On("GetUsers",
		mock.Anything,
		mock.AnythingOfType("Domain"),
	).Return(userlistDomain, nil).Once()

	t.Run("Test Case 1 | Valid GetUsers", func(t *testing.T) {
		_, err := userService.GetUsers(context.Background(), users.Domain{})
		assert.Nil(t, err)
	})

}

func TestGetUserById(t *testing.T) {
	setup()
	userRepository.On("GetUserById",
		mock.Anything,
		mock.AnythingOfType("Domain"),
	).Return(userDomain, nil).Once()

	t.Run("Test Case 1 | Valid GetUserById", func(t *testing.T) {
		_, err := userService.GetUserById(context.Background(), users.Domain{
			Id: 1,
		})
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Invalid GetUserById - No Id ", func(t *testing.T) {
		_, err := userService.GetUserById(context.Background(), users.Domain{
			Id: 0,
		})
		assert.NotNil(t, err)
	})

}
