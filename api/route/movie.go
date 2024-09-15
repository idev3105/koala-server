package route

import (
	"github.com/labstack/echo/v4"
	moviehandler "org.idev.koala/backend/api/handler/movie"
	"org.idev.koala/backend/app"
)

func NewMovieRouter(e *echo.Group, appCtx *app.AppContext) {
	handler := moviehandler.NewMovieHandler(appCtx)
	g := e.Group("/movies")
	{
		g.POST("", handler.CreateMovie())
		g.PUT("/:id", handler.UpdateMovie())
		g.GET("/:id/votes/stream", handler.StreamMovieVotes())
		g.PUT("/:id/votes", handler.VoteMovie())
	}
}
