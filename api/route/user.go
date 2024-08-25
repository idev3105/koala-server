package route

import (
	"github.com/labstack/echo/v4"
	userhandler "org.idev.koala/backend/api/handler/user"
	"org.idev.koala/backend/app"
)

func NewUserRouter(e *echo.Group, appCtx *app.AppContext) {
	handler := userhandler.NewUserHandler(appCtx)
	g := e.Group("/users")
	{
		g.POST("", handler.CreateUser())
		g.GET("/:id", handler.GetUserByUserId())
		g.POST("/exists", handler.ExistsUserByIdToken())
	}
}
