package routes

import (
	"myuseek/controller/users"

	"github.com/labstack/echo"
)

type ControllerList struct {
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.POST("users/register", cl.UserController.Register)
	e.POST("users/login", cl.UserController.Login)
}
