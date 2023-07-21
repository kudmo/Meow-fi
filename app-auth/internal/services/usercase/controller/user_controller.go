package controller

import (
	"Meow-fi_app-auth/internal/auth"
	"Meow-fi_app-auth/internal/config"
	"Meow-fi_app-auth/internal/models"
	"Meow-fi_app-auth/internal/services/usercase/repo"

	"github.com/labstack/echo/v4"
)

type UserInteractor struct {
	UserRepository repo.UserRepository
}

func (interactor *UserInteractor) CheckAuth(email, password string) (int, error) {
	user, err := interactor.UserRepository.SelectByEmail(email)

	if err != nil || user.Email != email || user.Password != auth.HashPassword(password, user.Salt, config.LocalSalt) {
		return 0, echo.ErrUnauthorized
	}

	return user.Id, nil
}

func (interactor *UserInteractor) Add(u models.User) error {
	return interactor.UserRepository.Store(u)
}
func (interactor *UserInteractor) UpdateRefreshToken(userId int, refreshId string) error {
	return interactor.UserRepository.UpdateRefreshToken(userId, refreshId)
}
func (interactor *UserInteractor) GetRefreshToken(userId int) (string, error) {
	return interactor.UserRepository.GetRefreshToken(userId)
}
func (interactor *UserInteractor) GetUserByEmail(login string) (models.User, error) {
	return interactor.UserRepository.SelectByEmail(login)
}
