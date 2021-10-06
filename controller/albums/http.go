package albums

import (
	"fmt"
	"myuseek/business/albums"
	controllers "myuseek/controller"
	"myuseek/controller/albums/requests"
	"myuseek/controller/albums/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AlbumController struct {
	AlbumUseCase albums.Usecase
}

func NewAlbumController(albumUseCase albums.Usecase) *AlbumController {
	return &AlbumController{
		AlbumUseCase: albumUseCase}
}

func (albumController AlbumController) Add(c echo.Context) error {
	fmt.Println("Add album")
	albumAdd := requests.AlbumAdd{}
	c.Bind(&albumAdd)

	ctx := c.Request().Context()
	album, error := albumController.AlbumUseCase.Add(ctx, albumAdd.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(album))
}

func (albumController AlbumController) GetAlbums(c echo.Context) error {
	fmt.Println("Get Albums")
	search := c.QueryParam("search")
	albumSearch := requests.AlbumSearch{}
	albumSearch.Title = search

	ctx := c.Request().Context()
	albumlistdomain, error := albumController.AlbumUseCase.GetAlbums(ctx, albumSearch.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromListDomainToGetAlbums(albumlistdomain))
}

func (albumController AlbumController) GetAlbumById(c echo.Context) error {
	fmt.Println("Get Album by ID")
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	albumByID := requests.AlbumByID{}
	albumByID.Id = id
	ctx := c.Request().Context()
	albumdomain, error := albumController.AlbumUseCase.GetAlbumById(ctx, albumByID.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomainToGetAlbumById(albumdomain))
}
