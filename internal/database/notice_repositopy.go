package database

import (
	"Meow-fi/internal/database/interfaces"
	"Meow-fi/internal/models"

	"gorm.io/gorm"
)

type NoticeRepository struct {
	interfaces.SqlHandler
}

func (db *NoticeRepository) Store(notice models.Notice) error {
	return db.Create(&notice)
}
func (db *NoticeRepository) Select() []models.Notice {
	var notices []models.Notice
	db.FindAll(&notices)
	return notices
}
func (db *NoticeRepository) UpdateNotice(notice models.Notice) error {
	return db.Update(&notice)
}
func (db *NoticeRepository) SelectById(id string) (models.Notice, error) {
	var notice models.Notice
	res := db.Preload("Client").Where("id = ?", id).Find(&notice)
	if res.Error != nil {
		return notice, res.Error
	}
	if res.RowsAffected == 0 {
		return notice, gorm.ErrRecordNotFound
	}
	return notice, nil
}
func (db *NoticeRepository) Delete(id string) error {
	var notices []models.Notice
	return db.DeleteById(&notices, id)
}
