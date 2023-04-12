package controller

import (
	"Meow-fi/internal/auth"
	"Meow-fi/internal/config"
	"Meow-fi/internal/models"
	"Meow-fi/internal/services/usercase/repo"

	"github.com/labstack/echo/v4"
)

type UserInteractor struct {
	UserRepository repo.UserRepository
}

func (interactor *UserInteractor) CheckAuth(login, password string) (int, error) {
	user, err := interactor.UserRepository.SelectByLogin(login)

	if err != nil || user.Login != login || user.Password != auth.HashPass(password, user.Salt, config.LocalSalt) {
		return 0, echo.ErrUnauthorized
	}

	return user.UserId, nil
}

func (interactor *UserInteractor) Add(u models.User) error {
	return interactor.UserRepository.Store(u)
}

func (interactor *UserInteractor) GetAllUsers() []models.User {
	return interactor.UserRepository.Select()
}
func (interactor *UserInteractor) GetUserByLogin(login string) (models.User, error) {
	return interactor.UserRepository.SelectByLogin(login)
}
func (interactor *UserInteractor) GetUserById(id int) (models.User, error) {
	return interactor.UserRepository.SelectById(id)
}
func (interactor *UserInteractor) Delete(id int) error {
	return interactor.UserRepository.Delete(id)
}
