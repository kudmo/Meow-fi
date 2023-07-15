package transport

import (
	"Meow-fi_app-auth/internal/auth"
	"Meow-fi_app-auth/internal/config"

	"Meow-fi_app-auth/internal/database"
	controllers "Meow-fi_app-auth/internal/services"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	userController := controllers.NewUserController(database.NewSqlHandler())

	userGroup := e.Group("/users")
	userGroup.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JWTClaims)
		},
		Skipper: func(c echo.Context) bool {
			return (c.Request().URL.Path == "/users/login" ||
				c.Request().URL.Path == "/users/relogin" ||
				c.Request().URL.Path == "/users/registrate")
		},
		SigningKey: []byte(config.SecretKeyJWT),
	}))
	userGroup.POST("/login", userController.Login)
	userGroup.POST("/relogin", userController.RefreshJWT)
	userGroup.POST("/registrate", userController.Registrate)
	userGroup.POST("/logout", userController.Logout)

	e.Logger.Fatal(e.Start(config.ServerPort))
}
