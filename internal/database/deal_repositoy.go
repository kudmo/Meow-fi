package database

import (
	"Meow-fi/internal/database/interfaces"
	"Meow-fi/internal/models"

	"gorm.io/gorm"
)

type DealRepository struct {
	interfaces.SqlHandler
}

func (db *DealRepository) Store(deal models.Deal) error {
	return db.Create(&deal)
}
func (db *DealRepository) Select() []models.Deal {
	var deals []models.Deal
	db.FindAll(&deals)
	return deals
}
func (db *DealRepository) UpdateDeal(deal models.Deal) error {
	return db.Update(deal)
}
func (db *DealRepository) SelectById(PerformerId, NoticeId int) (models.Deal, error) {
	var task models.Deal
	err := db.Where("performer_id = ?", PerformerId).Where("notice_id = ?", NoticeId).Find(&task)
	if err.RowsAffected == 0 {
		return task, gorm.ErrRecordNotFound
	}
	return task, err.Error
}
func (db *DealRepository) GetAllPerformerDeals(PerformerId int) ([]models.Deal, error) {
	var deals []models.Deal
	err := db.Where("performer_id = ?", PerformerId).Preload("Notice").Find(&deals).Error
	return deals, err
}
func (db *DealRepository) GetAllNoticeDeals(NoticeId int) ([]models.Deal, error) {
	var deals []models.Deal
	err := db.Where("notice_id = ?", NoticeId).Preload("Performer").Find(&deals).Error
	return deals, err
}
func (db *DealRepository) Delete(PerformerId, NoticeId int) error {
	var deals []models.Deal
	return db.Where("performer_id = ?", PerformerId).Where("notice_id = ?", NoticeId).Delete(&deals).Error
}
