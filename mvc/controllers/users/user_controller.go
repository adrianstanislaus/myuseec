package user_controllers

import (
	"myuseek/configs"
	"myuseek/models/response"
	"myuseek/models/users"
	"net/http"

	"github.com/labstack/echo"
)

func UserRegisterController(c echo.Context) error {
	var userRegister users.UserRegister
	c.Bind(&userRegister)

	// validasi
	if userRegister.FirstName == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Firstname empty",
			Data:    nil,
		})
	}
	if userRegister.LastName == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Lasstname empty",
			Data:    nil,
		})
	}
	if userRegister.Username == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Username empty",
			Data:    nil,
		})
	}
	if userRegister.Email == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Email empty",
			Data:    nil,
		})
	}
	if userRegister.Password == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Password empty",
			Data:    nil,
		})
	}
	if userRegister.Subscription_type == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Subscription type empty",
			Data:    nil,
		})
	}
	// etc

	var userDB users.User
	userDB.FirstName = userRegister.FirstName
	userDB.LastName = userRegister.LastName
	userDB.Email = userRegister.Email
	userDB.Password = userRegister.Password
	userDB.Username = userRegister.Username
	userDB.Bio = userRegister.Bio
	userDB.Profile_pic = userRegister.Profile_pic
	userDB.Subscription_type = userRegister.Subscription_type

	result := configs.DB.Create(&userDB)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error ketika input data user ke DB",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Berhasil register",
		Data:    userDB,
	})
}
