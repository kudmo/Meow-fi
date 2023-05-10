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
	//e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}))
	e.Use(middleware.Recover())

	noticeController := controllers.NewNoticeController(database.NewSqlHandler())
	userController := controllers.NewUserController(database.NewSqlHandler())
	fileController := controllers.NewMaterialController(database.NewSqlHandler())

	noticeHandler := &handlers.NoticeHandler{Controller: *noticeController}

	userGroup := e.Group("/users")
	userGroup.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JWTClaims)
		},
		Skipper: func(c echo.Context) bool {
			return !(c.Request().URL.Path == "/users/logout" || c.Request().URL.Path == "/users/delete")
		},
		SigningKey: []byte(config.SecretKeyJWT),
	}))
	userGroup.POST("/login", userController.Login)
	userGroup.POST("/logout", userController.Logout)
	userGroup.POST("/delete", userController.Delete)
	userGroup.POST("/relogin", userController.RefreshJWT)
	userGroup.POST("/registrate", userController.Registrate)
	userGroup.GET("/", userController.GetAllUsers)

	noticeGroup := e.Group("/notices")
	noticeGroup.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JWTClaims)
		},
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path == "/notices/"
		},
		SigningKey: []byte(config.SecretKeyJWT),
	}))
	noticeGroup.GET("/", noticeHandler.SelectWithFilter)
	noticeGroup.GET("/:id", noticeHandler.GetNoticeInfo)
	noticeGroup.PUT("/:id", noticeHandler.UpdateNotice)
	noticeGroup.DELETE("/:id", noticeHandler.DeleteNotice)
	noticeGroup.POST("/", noticeHandler.CreateNotice)
	noticeGroup.POST("/:id/deals", noticeHandler.AddResponse)
	noticeGroup.DELETE("/:id/deals", noticeHandler.DeleteDeal)
	noticeGroup.PUT("/:notice_id/deals/:performer_id", noticeHandler.ApproveResponse)

	fileGroup := e.Group("/files")
	fileGroup.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JWTClaims)
		},
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path != "/files/upload"
		},
		SigningKey: []byte(config.SecretKeyJWT),
	}))
	fileGroup.POST("/upload", fileController.Upload)
	fileGroup.GET("/", fileController.SelectWithFilter)
	fileGroup.GET("/:id/download", fileController.Download)

	e.Logger.Fatal(e.Start(config.ServerPort))
}
