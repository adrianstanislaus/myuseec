package routes

import (
	"myuseek/controller/users"

	"github.com/labstack/echo"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig      middleware.JWTConfig
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.POST("users/register", cl.UserController.Register)
	e.POST("users/login", cl.UserController.Login)
	//e.GET("users", cl.UserController.GetUsers, middleware.JWTWithConfig(cl.JwtConfig))
}
