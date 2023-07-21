package services

import (
	"Meow-fi_app-back/internal/auth"
	"Meow-fi_app-back/internal/database"
	"Meow-fi_app-back/internal/database/interfaces"
	"Meow-fi_app-back/internal/models"
	"Meow-fi_app-back/internal/services/usercase/controller"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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
func (controller *UserController) Update(c echo.Context) error {
	userId := auth.TokenGetUserId(c)

	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	user.Id = userId
	err := controller.Interactor.Add(user)
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.String(http.StatusCreated, "updated")
}

func (controller *UserController) GetAllUsers(c echo.Context) error {
	users := controller.Interactor.GetAllUsers()
	return c.JSON(http.StatusOK, users)
}
func (controller *UserController) GetUserInfo(c echo.Context) error {
	// type UserInfo struct {
	// 	Login string `json:"login"`
	// 	Email string `json:"email"`
	// }
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	user, err := controller.Interactor.GetUserById(id)
	if err == gorm.ErrRecordNotFound {
		return c.String(http.StatusNotFound, "no user")
	}
	return c.JSON(http.StatusOK, user)
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
