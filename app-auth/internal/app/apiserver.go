package app

import (
	"Meow-fi_app-auth/internal/config"
	"Meow-fi_app-auth/internal/models"
	"Meow-fi_app-auth/internal/transport"
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	state bool
}

func (server *Server) Start() {
	dbinit()
	transport.Init()
	e := echo.New()
	e.Logger.Fatal(e.Start(config.ServerPort))
}

func dbinit() {

	db, err := gorm.Open(postgres.Open(config.DatabaseUrl), &gorm.Config{})

	err = db.Migrator().CreateTable(models.User{})
	if err != nil {
		fmt.Print("User already exists")
	}
}
