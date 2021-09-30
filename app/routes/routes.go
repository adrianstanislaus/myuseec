package routes

import (
	"myuseek/controller/artists"
	"myuseek/controller/playlists"
	"myuseek/controller/songs"
	"myuseek/controller/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig          middleware.JWTConfig
	UserController     users.UserController
	ArtistController   artists.ArtistController
	SongController     songs.SongController
	PlaylistController playlists.PlaylistController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.POST("users/register", cl.UserController.Register)
	e.POST("users/login", cl.UserController.Login)
	e.GET("users", cl.UserController.GetUsers, middleware.JWTWithConfig(cl.JwtConfig))
	e.POST("artists/register", cl.ArtistController.Register, middleware.JWTWithConfig(cl.JwtConfig))
	e.GET("artists", cl.ArtistController.GetArtists, middleware.JWTWithConfig(cl.JwtConfig))
	e.POST("songs/add", cl.SongController.Add, middleware.JWTWithConfig(cl.JwtConfig))
	e.GET("songs", cl.SongController.GetSongs, middleware.JWTWithConfig(cl.JwtConfig))
	e.GET("songs/:id", cl.SongController.GetSongById, middleware.JWTWithConfig(cl.JwtConfig))
	e.GET("songs/:id/lyrics", cl.SongController.GetSongLyrics, middleware.JWTWithConfig(cl.JwtConfig))
	e.POST("playlists/create", cl.PlaylistController.Create, middleware.JWTWithConfig(cl.JwtConfig))
	e.GET("playlists", cl.PlaylistController.GetPlaylists, middleware.JWTWithConfig(cl.JwtConfig))
	e.GET("playlists/:id", cl.PlaylistController.GetbyID, middleware.JWTWithConfig(cl.JwtConfig))
	e.POST("playlists/:id/addsong", cl.PlaylistController.AddSong, middleware.JWTWithConfig(cl.JwtConfig))
}
