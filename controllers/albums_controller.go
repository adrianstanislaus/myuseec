package controllers

import (
	"myuseek/configs"
	"myuseek/models/albums"
	"myuseek/models/response"
	"net/http"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func GetAlbumsController(c echo.Context) error {
	albums := []albums.Album{}

	result := configs.DB.Find(&albums)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, response.BaseResponse{
				Code:    http.StatusInternalServerError,
				Message: "Error ketika input mendapatkan data user dari DB",
				Data:    nil,
			})
		}
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success retrieve albums data",
		Data:    albums,
	})
}
