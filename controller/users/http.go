package users

import (
	"fmt"
	"myuseek/business/users"
	controllers "myuseek/controller"
	"myuseek/controller/users/requests"
	"myuseek/controller/users/responses"
	"net/http"

	"github.com/labstack/echo"
)

type UserController struct {
	UserUseCase users.Usecase
}

func NewUserController(userUseCase users.Usecase) *UserController {
	return &UserController{
		UserUseCase: userUseCase}
}

func (userController UserController) Register(c echo.Context) error {
	fmt.Println("Register")
	userRegister := requests.UserRegister{}
	c.Bind(&userRegister)

	ctx := c.Request().Context()
	user, error := userController.UserUseCase.Register(ctx, userRegister.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}

func (userController UserController) Login(c echo.Context) error {
	fmt.Println("Login")
	userLogin := requests.UserLogin{}
	c.Bind(&userLogin)

	ctx := c.Request().Context()
	user, error := userController.UserUseCase.Login(ctx, userLogin.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}
func (userController UserController) GetUsers(c echo.Context) error {
	fmt.Println("GetUsers")

	ctx := c.Request().Context()
	userlistdomain, error := userController.UserUseCase.GetUsers(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromListDomain(userlistdomain))
}
