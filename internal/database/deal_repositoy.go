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
func (db *DealRepository) ApproveDeal(performerId, noticeId int) error {
	deal := models.Deal{}
	err := db.Where("performer_id = ?", performerId).Where("notice_id = ?", noticeId).Find(&deal)
	if err.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	deal.Approved = true
	return db.Update(deal)
}
func (db *DealRepository) SelectById(performerId, noticeId int) (models.Deal, error) {
	var task models.Deal
	err := db.Where("performer_id = ?", performerId).Where("notice_id = ?", noticeId).Find(&task)
	if err.RowsAffected == 0 {
		return task, gorm.ErrRecordNotFound
	}
	return task, err.Error
}
func (db *DealRepository) GetAllPerformerDeals(performerId int) ([]models.Deal, error) {
	var deals []models.Deal
	err := db.Where("performer_id = ?", performerId).Preload("Notice").Find(&deals).Error
	return deals, err
}
func (db *DealRepository) GetAllNoticeDeals(noticeId int) ([]models.Deal, error) {
	var deals []models.Deal
	err := db.Where("notice_id = ?", noticeId).Preload("Performer").Find(&deals).Error
	return deals, err
}
func (db *DealRepository) Delete(performerId, noticeId int) error {
	var deals []models.Deal
	res := db.Where("performer_id = ?", performerId).Where("notice_id = ?", noticeId)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Delete(&deals).Error
}
