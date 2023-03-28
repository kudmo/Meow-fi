package repo

import "Meow-fi/internal/models"

type UserRepository interface {
	Store(models.User) error
	Select() []models.User
	SelectById(id string) (models.User, error)
	SelectByLogin(login string) (models.User, error)
	Delete(id string) error
}
