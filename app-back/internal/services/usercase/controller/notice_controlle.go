package controller

import (
	"Meow-fi_app-back/internal/database"
	"Meow-fi_app-back/internal/models"
	"Meow-fi_app-back/internal/services/usercase/repo"
)

type NoticeInteractor struct {
	NoticeRepository repo.NoticeRepository
}

func (interactor *NoticeInteractor) Add(t models.Notice) error {
	return interactor.NoticeRepository.Store(t)
}

func (interactor *NoticeInteractor) CheckClient(userId, noticeId int) (bool, error) {
	return interactor.NoticeRepository.CheckClient(userId, noticeId)
}

func (interactor *NoticeInteractor) UpdateNotice(t models.Notice) error {
	return interactor.NoticeRepository.UpdateNotice(t)
}
func (interactor *NoticeInteractor) GetAllNotices() []models.Notice {
	return interactor.NoticeRepository.Select()
}
func (interactor *NoticeInteractor) GetNotice(id int) (models.Notice, error) {
	return interactor.NoticeRepository.SelectById(id)
}

func (interactor *NoticeInteractor) GetNoticeInfoFull(id int) (string, error) {
	notice, err := interactor.NoticeRepository.SelectById(id)
	if err != nil {
		return "", err
	}
	res := "Notice: \"" + notice.Containing +
		"\" (avaliable since " + notice.TimeAvaliable.Format("02-Jan-2006") + ")" +
		"created at " + notice.CreatedAt.Format("02-Jan-2006")

	return res, nil
}
func (interactor *NoticeInteractor) GetNoticeInfoShort(id int) (string, error) {
	notice, err := interactor.NoticeRepository.SelectById(id)
	if err != nil {
		return "", err
	}
	res := "Notice: \"" + notice.Containing +
		"\" (avaliable since " + notice.TimeAvaliable.Format("02-Jan-2006") + ")"

	return res, nil
}

func (interactor *NoticeInteractor) Delete(id int) error {
	return interactor.NoticeRepository.Delete(id)
}

func (interactor *NoticeInteractor) SelectWithFilter(filter database.SelectOptions) ([]models.Notice, error) {
	return interactor.NoticeRepository.SelectWithFilter(filter)
}
