package controller

import (
	"Meow-fi_app-back/internal/models"
	"Meow-fi_app-back/internal/services/usercase/repo"
)

type DealInteractor struct {
	DealRepository repo.DealRepository
}

func (interactor *DealInteractor) Add(t models.Deal) error {
	return interactor.DealRepository.Store(t)
}
func (interactor *DealInteractor) ApproveDeal(performerId, noticeId int) error {
	return interactor.DealRepository.ApproveDeal(performerId, noticeId)
}
func (interactor *DealInteractor) GetAllPerformerDeals(performerId int) ([]models.Deal, error) {
	return interactor.DealRepository.GetAllPerformerDeals(performerId)
}
func (interactor *DealInteractor) GetAllNoticeDeals(noticeId int) ([]models.Deal, error) {
	return interactor.DealRepository.GetAllNoticeDeals(noticeId)
}
func (interactor *DealInteractor) GetDeal(performerId, noticeId int) (models.Deal, error) {
	return interactor.DealRepository.SelectById(performerId, noticeId)
}
func (interactor *DealInteractor) Delete(performerId, noticeId int) error {
	return interactor.DealRepository.Delete(performerId, noticeId)
}
