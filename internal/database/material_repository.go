package database

import (
	"Meow-fi/internal/database/interfaces"
	"Meow-fi/internal/models"

	"gorm.io/gorm"
)

type MaterialRepository struct {
	interfaces.SqlHandler
}

func (db *MaterialRepository) Store(material models.Material) error {
	return db.Create(&material)
}
func (db *MaterialRepository) Select() []models.Material {
	var materials []models.Material
	db.FindAll(&materials)
	return materials
}
func (db *MaterialRepository) SelectWithFilter(filter SelectOptions) ([]models.Material, error) {
	var materials []models.Material
	var category models.Category
	var categoties *gorm.DB
	var res *gorm.DB
	res = nil
	if filter.categories != 0 {
		categoties = db.Where("category_id = ?", filter.categories).Take(&category)
		if categoties.Error != nil {
			return nil, categoties.Error
		}
		categoties = db.Table("categories").Where("left_key >= ? AND right_key <= ?", category.LeftKey, category.RightKey).Select("category_id")
		res = db.Where("category IN (?)", categoties)
		if res.Error != nil {
			return nil, res.Error
		}
	}
	if res != nil {
		res = res.Find(&materials)
		return materials, res.Error

	} else {
		err := db.FindAll(&materials)
		return materials, err
	}
}
func (db *MaterialRepository) ReadMaterialPath(id int) (string, error) {
	var material models.Material
	err := db.Where("id = ?", id).First(&material).Error
	return material.Path, err
}
func (db *MaterialRepository) SelectById(id int) (models.Material, error) {
	var material models.Material
	err := db.Where("id = ?", id).First(&material).Error
	return material, err
}
func (db *MaterialRepository) Delete(id int) error {
	var materials []models.Material
	return db.Where("id = ?", id).Delete(&materials).Error
}
