package app

import (
	"Meow-fi_app-back/internal/config"
	"Meow-fi_app-back/internal/models"
	"Meow-fi_app-back/internal/transport"
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
	err = db.Migrator().CreateTable(models.Notice{})
	if err != nil {
		fmt.Print("Notice already exists")
	}
	err = db.Migrator().CreateTable(models.Deal{})
	if err != nil {
		fmt.Print("Deal already exists")
	}
	err = db.Migrator().CreateTable(models.Category{})
	if err != nil {
		fmt.Print("Category already exists")
	}
	err = db.Migrator().CreateTable(models.Material{})
	if err != nil {
		fmt.Print("Material already exists")
	}
}
