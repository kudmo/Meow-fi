package controller

import (
	"Meow-fi_app-back/internal/models"
	"Meow-fi_app-back/internal/services/usercase/repo"
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

func (interactor *UserInteractor) GetUserById(id int) (models.User, error) {
	return interactor.UserRepository.SelectById(id)
}
func (interactor *UserInteractor) Delete(id int) error {
	return interactor.UserRepository.Delete(id)
}
