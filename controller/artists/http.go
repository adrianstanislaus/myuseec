package artists

import (
	"fmt"
	"myuseek/business/artists"
	controllers "myuseek/controller"
	"myuseek/controller/artists/requests"
	"myuseek/controller/artists/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ArtistController struct {
	ArtistUseCase artists.Usecase
}

func NewArtistController(artistUseCase artists.Usecase) *ArtistController {
	return &ArtistController{
		ArtistUseCase: artistUseCase}
}

func (artistController ArtistController) Register(c echo.Context) error {
	fmt.Println("Register Artist")
	artistRegister := requests.ArtistRegister{}
	c.Bind(&artistRegister)

	ctx := c.Request().Context()
	artist, error := artistController.ArtistUseCase.Register(ctx, artistRegister.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(artist))
}

func (artistController ArtistController) GetArtists(c echo.Context) error {
	fmt.Println("GetArtists")

	ctx := c.Request().Context()
	artistlistdomain, error := artistController.ArtistUseCase.GetArtists(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromListDomain(artistlistdomain))
}
