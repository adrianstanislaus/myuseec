package routes

import (
	"myuseek/controllers"

	"github.com/labstack/echo"
)

func NewRoute() *echo.Echo {
	e := echo.New()
	ev1 := e.Group("api/v1/")
	ev1.GET("albums", controllers.GetAlbumsController)
	return e
}
