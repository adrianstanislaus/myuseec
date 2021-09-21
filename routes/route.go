package routes

import (
	album_controllers "myuseek/controllers/albums"
	user_controllers "myuseek/controllers/users"

	"github.com/labstack/echo"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	ev1 := e.Group("api/v1/")
	ev1Album := ev1.Group("albums")
	ev1User := ev1.Group("users/register")
	ev1Album.GET("", album_controllers.GetAlbumsController)
	ev1User.POST("", user_controllers.UserRegisterController)
	return e
}
