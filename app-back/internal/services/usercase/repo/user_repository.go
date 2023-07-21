package repo

import "Meow-fi_app-back/internal/models"

type UserRepository interface {
	Store(models.User) error
	Select() []models.User
	SelectById(id int) (models.User, error)
	Delete(id int) error
}
