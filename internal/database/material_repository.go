package database

import (
	"Meow-fi/internal/config"
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
	if filter.Categories != 0 {
		categoties = db.Where("category_id = ?", filter.Categories).Take(&category)
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
		res = res.Order(filter.OrderBy)
	} else {
		res = db.Order(filter.OrderBy)
	}
	res = res.Offset(config.SizeNotionPage * filter.PageNumber).Limit(config.SizeNotionPage).Find(&materials)
	if res.Error != nil {
		return nil, res.Error
	}
	return materials, nil
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
