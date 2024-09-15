package middleware

import (
	"github.com/labstack/echo/v4"
	"org.idev.koala/backend/api/di"
	"org.idev.koala/backend/api/enum"
	"org.idev.koala/backend/app"
	"org.idev.koala/backend/common/logger"
	sqlc_generated "org.idev.koala/backend/generated/sqlc"
	tokenutil "org.idev.koala/backend/utils/token"
)

func AuthGuard(appCtx *app.AppContext) echo.MiddlewareFunc {

	log := logger.New("Middleware", "AuthGuard")

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			if ctx.Request().Header.Get("Authorization") == "" && appCtx.Config.Env == enum.Dev {
				ctx.Set("user_id", "user123")
				return next(ctx)
			}
			tokenRaw := tokenutil.GetTokenFromHeader(ctx.Request())

			isValid, err := tokenutil.Verify(ctx.Request().Context(), tokenRaw, appCtx.Config.JWKsUrl)
			if err != nil || !isValid {
				log.Errorf("Error verify token: %v", err)
				return echo.NewHTTPError(401, err.Error())
			}
			token, err := tokenutil.Parse(ctx.Request().Context(), tokenRaw, appCtx.Config.JWKsUrl)
			if err != nil {
				log.Errorf("Error verify token: %v", err)
				return echo.NewHTTPError(401, err.Error())
			}
			userId := token.Subject()

			userRepo := di.NewUserRepository(sqlc_generated.New(appCtx.Db), appCtx.Redis)

			user, err := userRepo.FindByUserId(ctx.Request().Context(), userId)
			if err != nil {
				log.Errorf("Error verify token: %v", err)
				return echo.NewHTTPError(401, err.Error())
			}
			ctx.Set("user", user)
			return next(ctx)
		}
	}
}
