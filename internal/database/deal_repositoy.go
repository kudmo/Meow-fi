package database

import (
	"Meow-fi/internal/database/interfaces"
	"Meow-fi/internal/models"
)

type DealRepository struct {
	interfaces.SqlHandler
}

func (db *DealRepository) Store(deal models.Deal) {
	db.Create(&deal)
}
func (db *DealRepository) Select() []models.Deal {
	var deals []models.Deal
	db.FindAll(&deals)
	return deals
}
func (db *DealRepository) UpdateDeal(deal models.Deal) {
	db.Update(deal)
}
func (db *DealRepository) SelectById(PerformerId string, NoticeId string) models.Deal {
	var task models.Deal
	db.Where("PerformerId = ?, NoticeId = ?", PerformerId, NoticeId).Find(&task)
	return task
}
func (db *DealRepository) GetDealInfo(PerformerId string, NoticeId string) models.Deal {
	var task models.Deal
	db.Preload("Client").Where("PerformerId = ?, NoticeId = ?", PerformerId, NoticeId).Find(&task)
	return task
}
func (db *DealRepository) Delete(PerformerId string, NoticeId string) {
	var deals []models.Deal
	db.DeleteByMultiId(&deals, PerformerId, NoticeId)
}
