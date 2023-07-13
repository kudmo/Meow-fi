package repo

import (
	"Meow-fi_app-back/internal/database"
	"Meow-fi_app-back/internal/models"
)

type MaterialRepository interface {
	Store(material models.Material) error
	Select() []models.Material
	SelectWithFilter(filter database.SelectOptions) ([]models.Material, error)
	ReadMaterialPath(id int) (string, error)
	SelectById(id int) (models.Material, error)
	Delete(id int) error
}
