package songs

import (
	"fmt"
	"myuseek/business/songs"
	controllers "myuseek/controller"
	"myuseek/controller/songs/requests"
	"myuseek/controller/songs/responses"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SongController struct {
	SongUseCase songs.Usecase
}

func NewSongController(songUseCase songs.Usecase) *SongController {
	return &SongController{
		SongUseCase: songUseCase}
}

func (songController SongController) Add(c echo.Context) error {
	fmt.Println("Add Song")
	songAdd := requests.SongAdd{}
	c.Bind(&songAdd)

	ctx := c.Request().Context()
	song, error := songController.SongUseCase.Add(ctx, songAdd.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(song))
}

func (songController SongController) GetSongs(c echo.Context) error {
	fmt.Println("GetSongs")

	ctx := c.Request().Context()
	songlistdomain, error := songController.SongUseCase.GetSongs(ctx)

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromListDomain(songlistdomain))
}

func (songController SongController) GetSongById(c echo.Context) error {
	fmt.Println("GetSong by ID")
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	songByID := requests.SongByID{}
	songByID.Id = id
	ctx := c.Request().Context()
	songdomain, error := songController.SongUseCase.GetSongById(ctx, songByID.ToDomain())

	if error != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, error)
	}

	return controllers.NewSuccesResponse(c, responses.FromDomain(songdomain))
}
