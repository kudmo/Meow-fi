package repo

import "Meow-fi/internal/models"

type DealRepository interface {
	Store(models.Deal) error
	UpdateDeal(models.Deal) error
	Select() []models.Deal
	SelectById(PerformerId string, NoticeId string) (models.Deal, error)
	GetDealInfo(Performerid string, NoticeId string) (models.Deal, error)
	Delete(Performerid string, NoticeId string) error
}
