package transport

import (
	"Meow-fi/internal/auth"
	"Meow-fi/internal/config"
	"Meow-fi/internal/handlers"

	"Meow-fi/internal/database"
	controllers "Meow-fi/internal/services"

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
	//dealController := controllers.NewDealController(database.NewSqlHandler())

	noticeHandler := &handlers.NoticeHandler{Controller: *noticeController}
	e.POST("login", userController.Login)
	e.GET("/users", userController.GetAllUsers)
	e.POST("/registrate", userController.Registrate)

	e.GET("/notices", noticeHandler.GetAllNotices)
	e.GET("/notices/:id",
		noticeHandler.GetNoticeInfo,
		echojwt.WithConfig(echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(auth.JwtCustomClaims)
			},
			SigningKey: []byte(config.SecretKeyJwt),
		}),
	)
	e.PUT("/notices/:id",
		noticeHandler.UpdateNotice,
		echojwt.WithConfig(echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(auth.JwtCustomClaims)
			},
			SigningKey: []byte(config.SecretKeyJwt),
		}),
	)
	e.DELETE("/notices/:id",
		noticeHandler.DeleteNotice,
		echojwt.WithConfig(echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(auth.JwtCustomClaims)
			},
			SigningKey: []byte(config.SecretKeyJwt),
		}),
	)
	e.POST("/notices",
		noticeHandler.Create,
		echojwt.WithConfig(echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(auth.JwtCustomClaims)
			},
			SigningKey: []byte(config.SecretKeyJwt),
		}),
	)

	e.POST("/notices/:id/deals",
		noticeHandler.AddResponse,
		echojwt.WithConfig(echojwt.Config{
			NewClaimsFunc: func(c echo.Context) jwt.Claims {
				return new(auth.JwtCustomClaims)
			},
			SigningKey: []byte(config.SecretKeyJwt),
		}),
	)

	e.Logger.Fatal(e.Start(config.ServerPort))
}
