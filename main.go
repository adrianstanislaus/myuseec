package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Album struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Album_type   string `json:"album_type"`
	About        string `json:"about"`
	Album_art    string `json:"album_art"`
	Release_date string `json:"release_date"`
}

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	e := echo.New()
	ev1 := e.Group("v1/")
	//routing
	ev1.GET("albums", GetAlbumController)
	//start server
	e.Start(":8000")
}

func GetAlbumController(c echo.Context) error {
	album := []Album{
		{1, "Ceritanya judul", "EP", "Ceritanya albumnya tentang apa gitu", "linkartalbum.com", "21 September 2021"},
		{2, "Ceritanya judul lagi", "Single", "Ceritanya albumnya tentang apa gitu", "linkartalbum2.com", "22 September 2021"},
	}

	return c.JSON(http.StatusOK, BaseResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    album,
	})
}
