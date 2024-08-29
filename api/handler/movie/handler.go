package moviehandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"org.idev.koala/backend/api/di"
	"org.idev.koala/backend/app"
	movieentity "org.idev.koala/backend/domain/movie/entity"
)

type MovieHandler struct {
	appCtx *app.AppContext
}

func NewMovieHandler(appCtx *app.AppContext) *MovieHandler {
	return &MovieHandler{
		appCtx: appCtx,
	}
}

// Create movie
// @Id CreateMovie
// @Summary Create new movie
// @Description
// @Tags movie
// @Accept json
// @Produce json
// @Param movie body CreateMovieRequest true "name"
// @Success 200 {object} MovieDto
// @Router /movies [post]
// @Security ApiKeyAuth
func (handler *MovieHandler) CreateMovie() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var data CreateMovieRequest
		if err := ctx.Bind(&data); err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}

		movieUseCase := di.NewMovieUseCase()
		movie := &movieentity.Movie{
			Name:         data.Name,
			Description:  data.Description,
			ThumbnailUrl: data.ThumbnailUrl,
		}

		createdMovie, err := movieUseCase.CreateMovie(ctx.Request().Context(), movie)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create movie"})
		}

		return ctx.JSON(http.StatusCreated, MovieDto{
			Id:          createdMovie.Id,
			Name:        createdMovie.Name,
			Description: createdMovie.Description,
		})
	}
}
