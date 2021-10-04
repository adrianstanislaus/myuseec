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
		Password:          "abc123",
		Bio:               "human being",
		Profile_pic:       "link.com",
		Subscription_type: "Premium",
		Token:             "123",
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}

func TestRegister(t *testing.T) {
	// t.Error("4 + 5 Harusnya 9")
	// assert.Equal(t, 2, 1)
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

	// t.Run("Test Case 2 | Invalid Email Empty", func(t *testing.T) {

	// 	_, err := userService.Login(context.Background(), users.Domain{
	// 		Email:    "",
	// 		Password: "123",
	// 	})
	// 	assert.NotNil(t, err)
	// })

	// t.Run("Test Case 3 | Invalid Password Empty", func(t *testing.T) {

	// 	_, err := userService.Login(context.Background(), users.Domain{
	// 		Email:    "alterra@gmail.com",
	// 		Password: "",
	// 	})
	// 	assert.NotNil(t, err)
	// })
}
