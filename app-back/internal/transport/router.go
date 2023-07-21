package transport

import (
	"Meow-fi_app-back/internal/auth"
	"Meow-fi_app-back/internal/config"
	"Meow-fi_app-back/internal/handlers"

	"Meow-fi_app-back/internal/database"
	controllers "Meow-fi_app-back/internal/services"

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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	noticeController := controllers.NewNoticeController(database.NewSqlHandler())
	userController := controllers.NewUserController(database.NewSqlHandler())
	fileController := controllers.NewMaterialController(database.NewSqlHandler())

	noticeHandler := &handlers.NoticeHandler{Controller: *noticeController}

	userGroup := e.Group("/users")
	userGroup.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JWTClaims)
		},
		SigningKey: []byte(config.SecretKeyJWT),
	}))
	userGroup.GET("/", userController.GetAllUsers)
	userGroup.GET("/:id", userController.GetAllUsers)
	userGroup.GET("/my/deals", noticeHandler.GetPerformerDeals)
	userGroup.POST("/update", userController.Update)
	userGroup.POST("/delete", userController.Delete)

	noticeGroup := e.Group("/notices")
	noticeGroup.Use(echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JWTClaims)
		},
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path == "/"
		},
		SigningKey: []byte(config.SecretKeyJWT),
	}))
	noticeGroup.GET("/", noticeHandler.SelectWithFilter)
	noticeGroup.GET("/:id", noticeHandler.GetNoticeInfo)
	noticeGroup.PUT("/:id", noticeHandler.UpdateNotice)
	noticeGroup.DELETE("/:id", noticeHandler.DeleteNotice)
	noticeGroup.POST("/new", noticeHandler.CreateNotice)
	noticeGroup.GET("/:id/deals", noticeHandler.GetNoticeDeals)
	noticeGroup.POST("/:id/deals", noticeHandler.AddResponse)
	noticeGroup.DELETE("/:id/deals", noticeHandler.DeleteDeal)
	noticeGroup.GET("/:notice_id/deals/:performer_id", noticeHandler.GetDealInfo)
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
