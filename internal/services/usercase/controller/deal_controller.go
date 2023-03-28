package controller

import (
	"Meow-fi/internal/models"
	"Meow-fi/internal/services/usercase/repo"
)

type DealInteractor struct {
	DealRepository repo.DealRepository
}

func (interactor *DealInteractor) Add(t models.Deal) error {
	return interactor.DealRepository.Store(t)
}
func (interactor *DealInteractor) UpdateDeal(t models.Deal) {
	interactor.DealRepository.UpdateDeal(t)
}
func (interactor *DealInteractor) GetAllDeals() []models.Deal {
	return interactor.DealRepository.Select()
}
func (interactor *DealInteractor) GetDeal(PerformerId string, NoticeId string) (models.Deal, error) {
	return interactor.DealRepository.SelectById(PerformerId, NoticeId)
}
func (interactor *DealInteractor) GetDealInfo(PerformerId string, NoticeId string) (models.Deal, error) {
	return interactor.DealRepository.GetDealInfo(PerformerId, NoticeId)
}
func (interactor *DealInteractor) Delete(PerformerId string, NoticeId string) error {
	return interactor.DealRepository.Delete(PerformerId, NoticeId)
}
