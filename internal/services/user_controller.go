package services

import (
	"Meow-fi/internal/auth"
	"Meow-fi/internal/config"
	"Meow-fi/internal/database"
	"Meow-fi/internal/database/interfaces"
	"Meow-fi/internal/models"
	"Meow-fi/internal/services/usercase/controller"
	"net/http"

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

func (controller *UserController) Login(c echo.Context) error {
	login := c.FormValue("login")
	password := c.FormValue("password")
	userId, err := controller.Interactor.CheckAuth(login, password)
	if err != nil {
		return c.String(http.StatusUnauthorized, "wrong login or password")
	}

	// Generate encoded token and send it as response.
	t, err := auth.CalculateToken(userId)
	if err != nil {
		return c.String(http.StatusInternalServerError, "something goes wrong "+err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
func (controller *UserController) Registrate(c echo.Context) error {
	login := c.FormValue("login")
	email := c.FormValue("email")
	password := c.FormValue("password")

	randomSalt := auth.RandSeq()
	hashedPass := auth.HashPass(password, randomSalt, config.LocalSalt)
	user := models.User{Login: login, Email: email, Password: hashedPass, Salt: randomSalt}

	if controller.Interactor.Add(user) != nil {
		return c.String(http.StatusBadRequest, "login or email already exist")
	}
	return c.String(http.StatusOK, "registrated")
}
func (controller *UserController) GetAllUsers(c echo.Context) error {
	users := controller.Interactor.GetAllUsers()
	return c.JSON(http.StatusOK, users)
}
func (controller *UserController) Delete(id int) error {
	err := controller.Interactor.Delete(id)
	return err
}
