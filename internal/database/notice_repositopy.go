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

func (db *NoticeRepository) CheckClient(userId, noticeId int) (bool, error) {
	var client_id models.Notice
	res := db.Where("id = ?", noticeId).Select("client_id").Find(&client_id)
	if res.RowsAffected == 0 {
		return false, gorm.ErrRecordNotFound
	}
	return client_id.ClientId == userId, res.Error
}

func (db *NoticeRepository) Select() []models.Notice {
	var notices []models.Notice
	db.FindAll(&notices)
	return notices
}
func (db *NoticeRepository) UpdateNotice(notice models.Notice) error {
	return db.Update(&notice)
}
func (db *NoticeRepository) SelectById(id int) (models.Notice, error) {
	var notice models.Notice
	res := db.Preload("Client").Where("id = ?", id).Find(&notice)
	if res.RowsAffected == 0 {
		return notice, gorm.ErrRecordNotFound
	}
	return notice, res.Error
}
func (db *NoticeRepository) Delete(id int) error {
	var notices []models.Notice
	return db.Where("id = ?", id).Delete(&notices).Error
}
