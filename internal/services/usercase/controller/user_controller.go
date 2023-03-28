package controller

import (
	"Meow-fi/internal/models"
	"Meow-fi/internal/services/usercase/repo"
)

type UserInteractor struct {
	UserRepository repo.UserRepository
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
func (interactor *UserInteractor) GetUserById(id string) (models.User, error) {
	return interactor.UserRepository.SelectById(id)
}
func (interactor *UserInteractor) Delete(id string) error {
	return interactor.UserRepository.Delete(id)
}
