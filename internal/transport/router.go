package transport

import (
	"Meow-fi/internal/auth"
	"Meow-fi/internal/config"
	"Meow-fi/internal/handlers"

	"Meow-fi/internal/database"
	controllers "Meow-fi/internal/services"
	"time"

	"net/http"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	noticeController := controllers.NewNoticeController(database.NewSqlHandler())
	userController := controllers.NewUserController(database.NewSqlHandler())

	noticeHandler := &handlers.NoticeHandler{Controller: *noticeController}

	e.POST("login", func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		// Throws unauthorized error
		user, err := userController.GetUserByLogin(username)
		if err != nil {
			return echo.ErrUnauthorized
		}
		if user.Login != username || user.Password != auth.HashPass(password, user.Salt, config.LocalSalt) {
			return echo.ErrUnauthorized
		}
		// Set custom claims
		claims := &auth.JwtCustomClaims{
			user.UserId,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(config.SecretKeyJwt))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	})

	e.GET("/notices", func(ctx echo.Context) error {
		tasks := noticeController.GetAllNotices()
		return ctx.JSON(http.StatusOK, tasks)
	})
	e.GET("/notices/:id",
		func(ctx echo.Context) error {
			return ctx.NoContent(http.StatusOK)
		},
		echojwt.WithConfig(echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(auth.JwtCustomClaims)
			},
			SigningKey: []byte(config.SecretKeyJwt),
			SuccessHandler: func(c echo.Context) {
				noticeHandler.GetNoticeInfo(c)
			},
			ErrorHandler: func(c echo.Context, err error) error {
				noticeHandler.GetNoticeInfoGuest(c)
				return nil
			},
		}))

	e.PUT("/notices/:id",
		func(ctx echo.Context) error {
			return ctx.NoContent(http.StatusOK)
		},
		echojwt.WithConfig(echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(auth.JwtCustomClaims)
			},
			SigningKey: []byte(config.SecretKeyJwt),
			SuccessHandler: func(c echo.Context) {
				noticeHandler.UpdateNotice(c)
			},
		}))

	e.DELETE("/notices/:id",
		func(ctx echo.Context) error {
			return ctx.NoContent(http.StatusOK)
		},
		echojwt.WithConfig(echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(auth.JwtCustomClaims)
			},
			SigningKey: []byte(config.SecretKeyJwt),
			SuccessHandler: func(c echo.Context) {
				noticeHandler.DeleteNotice(c)
			},
		}))

	e.POST("/notices", func(ctx echo.Context) error {
		noticeController.Create(ctx)
		return ctx.String(http.StatusOK, "created")
	})

	e.GET("/users", func(ctx echo.Context) error {
		users := userController.GetAllUsers()
		return ctx.JSON(http.StatusOK, users)
	})

	e.POST("/registrate", func(c echo.Context) error {
		err := userController.Create(c)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.String(http.StatusOK, "created")
	})

	e.Logger.Fatal(e.Start(config.ServerPort))
}
