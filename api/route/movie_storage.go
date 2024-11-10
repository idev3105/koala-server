package route

import (
	"github.com/labstack/echo/v4"
	moviehandler "org.idev.koala/backend/api/handler/movie"
	"org.idev.koala/backend/app"
)

func NewMovieStorageRouter(e *echo.Group, appCtx *app.AppContext) {
	handler := moviehandler.NewMovieHandler(appCtx)
	g := e.Group("/movies")
	{
		g.GET("/:id/:episodeId/playlist.m3u8", handler.GetEpisodeVideo())
		g.GET("/:id/:episodeId/:segmentId", handler.GetVideoSegment())
	}
}
