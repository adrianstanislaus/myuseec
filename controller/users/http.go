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
	user, error := userController.UserUseCase.Register(ctx, userRegister.FirstName, userRegister.LastName, userRegister.Username, userRegister.Email, userRegister.Password, userRegister.Bio, userRegister.Profile_pic, userRegister.Subscription_type)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(user))
}
