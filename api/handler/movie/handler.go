package moviehandler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"org.idev.koala/backend/api/di"
	"org.idev.koala/backend/app"
	"org.idev.koala/backend/component/kafka"
	movieentity "org.idev.koala/backend/domain/movie/entity"
	sqlc_generated "org.idev.koala/backend/generated/sqlc"
	"org.idev.koala/backend/utils"
)

type MovieHandler struct {
	appCtx *app.AppContext
}

func NewMovieHandler(appCtx *app.AppContext) *MovieHandler {
	return &MovieHandler{
		appCtx: appCtx,
	}
}

// CreateMovie handles the creation of a new movie
// @Id CreateMovie
// @Summary Create new movie
// @Description Create a new movie entry
// @Tags movie
// @Accept json
// @Produce json
// @Param movie body CreateMovieRequest true "Movie details"
// @Success 201 {object} MovieDto
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /movies [post]
// @Security ApiKeyAuth
func (handler *MovieHandler) CreateMovie() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var data CreateMovieRequest
		if err := ctx.Bind(&data); err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}

		movie := &movieentity.Movie{
			Name:         data.Name,
			Description:  data.Description,
			ThumbnailUrl: data.ThumbnailUrl,
		}

		tx, err := handler.appCtx.Db.BeginTx(ctx.Request().Context(), pgx.TxOptions{})
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to begin transaction"})
		}
		defer func() {
			if err != nil {
				tx.Rollback(ctx.Request().Context())
			}
		}()

		queries := sqlc_generated.New(handler.appCtx.Db)
		movieUseCase := di.NewMovieUseCase(queries)
		createdMovie, err := movieUseCase.CreateMovie(ctx.Request().Context(), movie)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create movie"})
		}

		if err = handler.sendCreateMovieMessage(createdMovie.Id); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send message to Kafka"})
		}

		if err = tx.Commit(ctx.Request().Context()); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to commit transaction"})
		}

		return ctx.JSON(http.StatusCreated, MovieDto{
			MovieID:      createdMovie.MovieID,
			Name:         createdMovie.Name,
			Description:  createdMovie.Description,
			ThumbnailUrl: createdMovie.ThumbnailUrl,
		})
	}
}

// UpdateMovie handles updating a movie's details
// @Id UpdateMovie
// @Summary Update movie details
// @Description Update the details of an existing movie
// @Tags movie
// @Accept json
// @Produce json
// @Param id path string true "Movie ID"
// @Param movie body UpdateMovieRequest true "Updated movie details"
// @Success 200 {object} MovieDto
// @Router /movies/{id} [put]
// @Security ApiKeyAuth
func (handler *MovieHandler) UpdateMovie() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		movieID, err := strconv.ParseInt(ctx.Param("id"), 10, 32)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid movie ID"})
		}

		var data UpdateMovieRequest
		if err := ctx.Bind(&data); err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		}

		updatedMovie := &movieentity.Movie{
			Id:           int32(movieID),
			Name:         data.Name,
			Description:  data.Description,
			ThumbnailUrl: data.ThumbnailUrl,
		}

		tx, err := handler.appCtx.Db.BeginTx(ctx.Request().Context(), pgx.TxOptions{})
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to begin transaction"})
		}
		defer func() {
			if err != nil {
				tx.Rollback(ctx.Request().Context())
			}
		}()

		queries := sqlc_generated.New(handler.appCtx.Db)
		movieUseCase := di.NewMovieUseCase(queries)
		result, err := movieUseCase.UpdateMovie(ctx.Request().Context(), updatedMovie)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update movie"})
		}

		if err = handler.sendUpdateMovieMessage(updatedMovie.Id); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send message to Kafka"})
		}

		if err = tx.Commit(ctx.Request().Context()); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to commit transaction"})
		}

		return ctx.JSON(http.StatusOK, MovieDto{
			MovieID:      result.MovieID,
			Name:         result.Name,
			Description:  result.Description,
			ThumbnailUrl: result.ThumbnailUrl,
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

		tx, err := handler.appCtx.Db.BeginTx(ctx.Request().Context(), pgx.TxOptions{})
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to begin transaction"})
		}
		defer func() {
			if err != nil {
				tx.Rollback(ctx.Request().Context())
			}
		}()

		queries := sqlc_generated.New(handler.appCtx.Db)
		movieUseCase := di.NewMovieUseCase(queries)
		updatedMovie, err := movieUseCase.VoteMovie(ctx.Request().Context(), movieID, userID, data.VoteType == "up")
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to vote for movie"})
		}

		if err = handler.sendVoteMovieMessage(MovieVotedMessage{
			MovieID:  updatedMovie.MovieID,
			UserID:   userID,
			VoteType: data.VoteType,
		}); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to send message to Kafka"})
		}

		if err = tx.Commit(ctx.Request().Context()); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to commit transaction"})
		}

		return ctx.JSON(http.StatusOK, MovieDto{
			MovieID:     updatedMovie.MovieID,
			Name:        updatedMovie.Name,
			Description: updatedMovie.Description,
		})
	}
}

func (handler *MovieHandler) sendCreateMovieMessage(movieID int32) error {
	message := kafka.Message{
		Key:   []byte("movie_created"),
		Value: []byte(fmt.Sprintf(`{"movie_id":%d}`, movieID)),
	}
	return handler.appCtx.KafkaProducer.SendSync([]kafka.Message{message})
}

func (handler *MovieHandler) sendUpdateMovieMessage(movieID int32) error {
	message := kafka.Message{
		Key:   []byte("movie_updated"),
		Value: []byte(fmt.Sprintf(`{"movie_id":%d}`, movieID)),
	}
	return handler.appCtx.KafkaProducer.SendSync([]kafka.Message{message})
}

// StreamMovieVotes handles SSE connections for real-time movie vote updates
// @Id StreamMovieVotes
// @Summary Stream movie vote updates
// @Description Establish an SSE connection to receive real-time movie vote updates
// @Tags movie
// @Produce text/event-stream
// @Param id path string true "Movie ID"
// @Success 200 {string} string "SSE stream established"
// @Router /movies/{id}/votes/stream [get]
// @Security ApiKeyAuth
func (handler *MovieHandler) StreamMovieVotes() echo.HandlerFunc {
	return func(c echo.Context) error {
		movieID := c.Param("id")

		utils.SetSSEHeaders(c.Response())
		c.Response().WriteHeader(http.StatusOK)

		voteChan := make(chan string)

		go handler.subscribeToMovieVotes(c.Request().Context(), movieID, voteChan)

		for {
			select {
			case vote := <-voteChan:
				fmt.Println(vote)
				if _, err := c.Response().Write([]byte(vote)); err != nil {
					return err
				}
				c.Response().Flush()
			case <-c.Request().Context().Done():
				close(voteChan)
				return nil
			case <-time.After(30 * time.Second):
				fmt.Fprintf(c.Response(), ": keepalive\n\n")
				c.Response().Flush()
			}
		}
	}
}

func (handler *MovieHandler) GetEpisodeVideo() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// Get movie ID and episode ID from the request parameters
		movieId := ctx.Param("id")
		episodeId := ctx.Param("episodeId")

		playlistKey := fmt.Sprintf("%s/%s/playlist.m3u8", movieId, episodeId)
		data, err := handler.appCtx.StorageCli.Get(ctx.Request().Context(), "movies", playlistKey)

		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}

		ctx.Response().Header().Set(echo.HeaderContentType, "application/vnd.apple.mpegurl")
		return ctx.Stream(http.StatusOK, "application/vnd.apple.mpegurl", bytes.NewReader(data))
	}
}

func (handler *MovieHandler) GetVideoSegment() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		movieId := ctx.Param("id")
		episodeId := ctx.Param("episodeId")
		segmentId := ctx.Param("segmentId")

		data, err := handler.appCtx.StorageCli.Get(ctx.Request().Context(), "movies", fmt.Sprintf("%s/%s/%s", movieId, episodeId, segmentId))

		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}

		ctx.Response().Header().Set(echo.HeaderContentType, "application/vnd.apple.mpegurl")
		return ctx.Stream(http.StatusOK, "application/vnd.apple.mpegurl", bytes.NewReader(data))
	}
}

func (handler *MovieHandler) sendVoteMovieMessage(msg MovieVotedMessage) error {
	jsonEvent, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	message := kafka.Message{
		Key:   []byte("movie_voted"),
		Value: jsonEvent,
		Topic: "movie_voted",
	}

	return handler.appCtx.KafkaProducer.SendSync([]kafka.Message{message})
}

func (handler *MovieHandler) subscribeToMovieVotes(ctx context.Context, movieID string, eventChan chan<- string) error {
	consumer, err := kafka.NewConsumer(handler.appCtx.Config.KafkaHost, handler.appCtx.Config.KafkaPort, "movie_voted")
	if err != nil {
		return err
	}

	err = consumer.Consume(ctx, func(message []byte) error {
		var event MovieVotedMessage
		if err := json.Unmarshal(message, &event); err != nil {
			return err
		}
		if event.MovieID == movieID {
			eventChan <- string(message)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
