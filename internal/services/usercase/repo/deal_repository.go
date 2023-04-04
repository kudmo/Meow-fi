package repo

import "Meow-fi/internal/models"

type DealRepository interface {
	Store(models.Deal) error
	UpdateDeal(models.Deal) error
	Select() []models.Deal
	SelectById(PerformerId, NoticeId int) (models.Deal, error)
	GetAllPerformerDeals(PerformerId int) ([]models.Deal, error)
	GetAllNoticeDeals(NoticeId int) ([]models.Deal, error)
	Delete(Performerid, NoticeId int) error
}
