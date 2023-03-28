package services

import (
	"Meow-fi/internal/auth"
	"Meow-fi/internal/config"
	"Meow-fi/internal/database"
	"Meow-fi/internal/database/interfaces"
	"Meow-fi/internal/models"
	"Meow-fi/internal/services/usercase/controller"
	"errors"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Interactor controller.UserInteractor
}

func NewUserController(sqlHandler interfaces.SqlHandler) *UserController {
	return &UserController{
		Interactor: controller.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c echo.Context) error {
	login := c.FormValue("login")
	password := c.FormValue("password")

	randomSalt := auth.RandSeq()
	hashedPass := auth.HashPass(password, randomSalt, config.LocalSalt)
	user := models.User{Login: login, Password: hashedPass, Salt: randomSalt}

	if controller.Interactor.Add(user) != nil {
		return errors.New("login already exist")
	}
	return nil
}
func (controller *UserController) GetAllUsers() []models.User {
	res := controller.Interactor.GetAllUsers()
	return res
}
func (controller *UserController) GetUserByLogin(login string) (models.User, error) {
	return controller.Interactor.GetUserByLogin(login)
}
func (controller *UserController) Delete(id string) {
	controller.Interactor.Delete(id)
}
