package repo

import "Meow-fi/internal/models"

type DealRepository interface {
	Store(models.Deal)
	UpdateDeal(models.Deal)
	Select() []models.Deal
	SelectById(PerformerId string, NoticeId string) models.Deal
	GetDealInfo(Performerid string, NoticeId string) models.Deal
	Delete(Performerid string, NoticeId string)
}
