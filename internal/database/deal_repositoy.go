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
func (db *DealRepository) SelectById(PerformerId string, NoticeId string) (models.Deal, error) {
	var task models.Deal
	err := db.Where("performer_id = ?", PerformerId).Where("notice_id = ?", NoticeId).Find(&task)
	if err.RowsAffected == 0 {
		return task, gorm.ErrRecordNotFound
	}
	return task, err.Error
}
func (db *DealRepository) GetDealInfo(PerformerId string, NoticeId string) (models.Deal, error) {
	var task models.Deal
	err := db.Preload("Performer").Where("performer_id = ?", PerformerId).Preload("Notice").Where("notice_id = ?", NoticeId).Find(&task)
	if err.RowsAffected == 0 {
		return task, gorm.ErrRecordNotFound
	}
	return task, err.Error
}
func (db *DealRepository) Delete(PerformerId string, NoticeId string) error {
	var deals []models.Deal
	return db.Where("performer_id = ?", PerformerId).Where("notice_id = ?", NoticeId).Delete(&deals).Error
}
