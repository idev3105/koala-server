package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"org.idev.koala/backend/api/di"
	"org.idev.koala/backend/app"
	sqlc_generated "org.idev.koala/backend/generated/sqlc"
	tokenutil "org.idev.koala/backend/utils/token"
)

// User gin handler
type UserHandler struct {
	appCtx *app.AppContext
}

func NewUserHandler(appCtx *app.AppContext) *UserHandler {
	return &UserHandler{appCtx: appCtx}
}

// Create new user
// @Id CreateUser
// @Summary Create new user
// @Description Create new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "username"
// @Success 200 {object} UserDto
// @Router /users [post]
// @Security ApiKeyAuth
func (u *UserHandler) CreateUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// bind body from request
		var data CreateUserRequest
		err := (&echo.DefaultBinder{}).BindBody(ctx, &data)
		if err != nil {
			panic(err)
		}

		token, err := tokenutil.Parse(ctx.Request().Context(), data.IdToken, u.appCtx.Config.JWKsUrl)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		userId := token.Subject()

		userUseCase := di.NewUserUseCase(sqlc_generated.New(u.appCtx.Db), u.appCtx.Redis)
		user, err := userUseCase.Create(ctx.Request().Context(), userId, data.Username)
		if err != nil {
			panic(err)
		}

		return ctx.JSON(http.StatusOK, UserDto{
			Id:       user.UserId,
			Username: user.Username,
		})
	}
}

// Get user by user id
// @Id GetUserByUserId
// @Summary Get user by user id
// @Description Get user by user id
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} UserDto
// @Router /users/{id} [get]
func (u *UserHandler) GetUserByUserId() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId := ctx.Param("id")

		userUseCase := di.NewUserUseCase(sqlc_generated.New(u.appCtx.Db), u.appCtx.Redis)
		user, err := userUseCase.FindByUserId(ctx.Request().Context(), userId)
		if err != nil {
			panic(err)
		}

		return ctx.JSON(http.StatusOK, UserDto{
			Id:       user.UserId,
			Username: user.Username,
		})
	}
}

// Check exist user by token
// @Id existsUserByIdToken
// @Summary Check exist user by id token
// @Description Check user exist
// @Tags user
// @Accept json
// @Produce json
// @Param idToken body CheckUserByIdTokenRequest true "idToken"
// @Success 200 {object} CheckUserByIdTokenResponse
// @Router /users/exists [post]
func (u *UserHandler) ExistsUserByIdToken() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var data CheckUserByIdTokenRequest
		err := (&echo.DefaultBinder{}).BindBody(ctx, &data)
		if err != nil {
			panic(err)
		}
		token, err := tokenutil.Parse(ctx.Request().Context(), data.IdToken, u.appCtx.Config.JWKsUrl)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		userId := token.Subject()

		userUseCase := di.NewUserUseCase(sqlc_generated.New(u.appCtx.Db), u.appCtx.RedisCli)
		exist, err := userUseCase.ExistsByUserId(ctx.Request().Context(), userId)
		if err != nil {
			panic(err)
		}

		return ctx.JSON(http.StatusOK, CheckUserByIdTokenResponse{Exist: exist})
	}
}
