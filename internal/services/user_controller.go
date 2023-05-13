package services

import (
	"Meow-fi/internal/auth"
	"Meow-fi/internal/config"
	"Meow-fi/internal/database"
	"Meow-fi/internal/database/interfaces"
	"Meow-fi/internal/models"
	"Meow-fi/internal/services/usercase/controller"
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Interactor      controller.UserInteractor
	TokenController auth.AuthController
}

func NewUserController(sqlHandler interfaces.SqlHandler) *UserController {
	return &UserController{
		Interactor: controller.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
		TokenController: auth.AuthController{},
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
	rtoken, rtoken_id, err := controller.TokenController.CalculateRefreshToken(userId)
	if err != nil {
		log.Println("error while generating tokens: " + err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	err = controller.Interactor.UpdateRefreshToken(userId, rtoken_id)
	if err != nil {
		log.Println("error while generating tokens: " + err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	atoken, _, err := controller.TokenController.CalculateAccessToken(userId, rtoken_id)
	if err != nil {
		log.Println("error while generating tokens: " + err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"access_token":  atoken,
		"refresh_token": rtoken,
	})
}
func (controller *UserController) Logout(c echo.Context) error {
	userId := auth.TokenGetUserId(c)
	err := controller.Interactor.UpdateRefreshToken(userId, "")
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.String(http.StatusOK, "logget out")
}
func (controller *UserController) RefreshJWT(c echo.Context) error {
	access, refresh, err := controller.TokenController.GetTokensFromContext(c)
	if errors.Is(err, jwt.ErrTokenExpired) {
		return c.NoContent(http.StatusUnauthorized)
	}
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	currId, err := controller.Interactor.UserRepository.GetRefreshToken(access.UserId)
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	if currId != refresh.TokenId {
		return c.NoContent(http.StatusUnauthorized)
	}
	newRefreshToken, newId, err := controller.TokenController.CalculateRefreshToken(access.UserId)
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	err = controller.Interactor.UpdateRefreshToken(access.UserId, newId)
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	newAccess, _, err := controller.TokenController.CalculateAccessToken(access.UserId, newId)
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"access_token":  newAccess,
		"refresh_token": newRefreshToken,
	})
}
func (controller *UserController) Registrate(c echo.Context) error {
	login := c.FormValue("login")
	email := c.FormValue("email")
	password := c.FormValue("password")

	randomSalt := auth.GenerateRandomSalt()
	hashedPass := auth.HashPassword(password, randomSalt, config.LocalSalt)
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
func (controller *UserController) Delete(c echo.Context) error {
	userId := auth.TokenGetUserId(c)
	err := controller.Interactor.Delete(userId)
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.String(http.StatusOK, "deleted")

}
