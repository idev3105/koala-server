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

// VoteMovie handles upvoting or downvoting a movie
// @Id VoteMovie
// @Summary Vote for a movie
// @Description Upvote or downvote a movie
// @Tags movie
// @Accept json
// @Produce json
// @Param id path string true "Movie ID"
// @Param vote body VoteMovieRequest true "Vote details"
// @Success 200 {object} MovieDto
// @Router /movies/{id}/vote [post]
// @Security ApiKeyAuth
func (handler *MovieHandler) VoteMovie() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		movieID := ctx.Param("id")
		userID := ctx.Get("user_id").(string)

		var data VoteMovieRequest
		if err := ctx.Bind(&data); err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}

		movieUseCase := di.NewMovieUseCase()
		updatedMovie, err := movieUseCase.VoteMovie(ctx.Request().Context(), movieID, userID, data.VoteType == "up")
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to vote for movie"})
		}

		return ctx.JSON(http.StatusOK, MovieDto{
			Id:          updatedMovie.Id,
			Name:        updatedMovie.Name,
			Description: updatedMovie.Description,
		})
	}
}
