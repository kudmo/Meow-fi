package repo

import "Meow-fi_app-auth/internal/models"

type UserRepository interface {
	Store(models.User) error
	UpdateRefreshToken(int, string) error
	GetRefreshToken(int) (string, error)
	Select() []models.User
	SelectById(id int) (models.User, error)
	SelectByEmail(email string) (models.User, error)
	Delete(id int) error
}
