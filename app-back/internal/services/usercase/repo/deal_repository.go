package repo

import "Meow-fi_app-back/internal/models"

type DealRepository interface {
	Store(models.Deal) error
	ApproveDeal(performerId, noticeId int) error
	Select() []models.Deal
	SelectById(performerId, noticeId int) (models.Deal, error)
	GetAllPerformerDeals(performerId int) ([]models.Deal, error)
	GetAllNoticeDeals(noticeId int) ([]models.Deal, error)
	Delete(performerId, noticeId int) error
}
