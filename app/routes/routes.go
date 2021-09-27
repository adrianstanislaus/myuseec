package routes

import (
	"myuseek/controller/artists"
	"myuseek/controller/songs"
	"myuseek/controller/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig        middleware.JWTConfig
	UserController   users.UserController
	ArtistController artists.ArtistController
	SongController   songs.SongController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.POST("users/register", cl.UserController.Register)
	e.POST("users/login", cl.UserController.Login)
	e.GET("users", cl.UserController.GetUsers, middleware.JWTWithConfig(cl.JwtConfig))
	e.POST("artists/register", cl.ArtistController.Register, middleware.JWTWithConfig(cl.JwtConfig))
	e.GET("artists", cl.ArtistController.GetArtists, middleware.JWTWithConfig(cl.JwtConfig))
	e.POST("songs/add", cl.SongController.Add, middleware.JWTWithConfig(cl.JwtConfig))
	e.GET("songs", cl.SongController.GetSongs, middleware.JWTWithConfig(cl.JwtConfig))
}
