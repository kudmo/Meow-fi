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
func (interactor *DealInteractor) UpdateDeal(t models.Deal) error {
	return interactor.DealRepository.UpdateDeal(t)
}
func (interactor *DealInteractor) GetAllPerformerDeals(PerformerId int) ([]models.Deal, error) {
	return interactor.DealRepository.GetAllPerformerDeals(PerformerId)
}
func (interactor *DealInteractor) GetAllNoticeDeals(NoticeId int) ([]models.Deal, error) {
	return interactor.DealRepository.GetAllNoticeDeals(NoticeId)
}
func (interactor *DealInteractor) GetDeal(PerformerId, NoticeId int) (models.Deal, error) {
	return interactor.DealRepository.SelectById(PerformerId, NoticeId)
}
func (interactor *DealInteractor) Delete(PerformerId, NoticeId int) error {
	return interactor.DealRepository.Delete(PerformerId, NoticeId)
}
