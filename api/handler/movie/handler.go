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
		err := (&echo.DefaultBinder{}).BindBody(ctx, &data)
		if err != nil {
			panic(err)
		}
		movieUseCase := di.NewMovieUseCase()
		movie := &movieentity.Movie{
			Name:         data.Name,
			Description:  data.Description,
			ThumbnailUrl: data.ThumbnailUrl,
		}

		movie, err = movieUseCase.CreateMovie(ctx.Request().Context(), movie)
		if err != nil {
			panic(err)
		}
		return ctx.JSON(http.StatusOK, MovieDto{
			Id:          movie.Id,
			Name:        movie.Name,
			Description: movie.Description,
		})
	}
}
