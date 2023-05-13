package services

import (
	"Meow-fi/internal/database"
	"Meow-fi/internal/database/interfaces"
	"Meow-fi/internal/models"
	"Meow-fi/internal/services/usercase/controller"
	"errors"
)

type NoticeController struct {
	noticeInteractor controller.NoticeInteractor
	dealInteractor   controller.DealInteractor
}

func NewNoticeController(sqlHandler interfaces.SqlHandler) *NoticeController {
	return &NoticeController{
		noticeInteractor: controller.NoticeInteractor{
			NoticeRepository: &database.NoticeRepository{
				SqlHandler: sqlHandler,
			},
		},
		dealInteractor: controller.DealInteractor{
			DealRepository: &database.DealRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *NoticeController) CheckClient(userId, noticeId int) (bool, error) {
	return controller.noticeInteractor.CheckClient(userId, noticeId)
}
func (controller *NoticeController) Create(idUser int, notice models.Notice) error {
	notice.ClientId = idUser
	err := controller.noticeInteractor.Add(notice)
	return err
}
func (controller *NoticeController) UpdateNotice(idUser int, noticeId int, notice models.Notice) error {
	notice_, err := controller.noticeInteractor.GetNotice(noticeId)
	if err != nil {
		return err
	}
	if notice_.ClientId != idUser {
		return errors.New("not owner")
	}
	notice.ClientId = idUser
	notice.Id = notice_.Id
	notice.TypeNotice = notice_.TypeNotice
	err = controller.noticeInteractor.UpdateNotice(notice)
	return err
}
func (controller *NoticeController) GetNotice(noticeId int) (models.Notice, error) {
	notice, err := controller.noticeInteractor.GetNotice(noticeId)
	return notice, err
}

// Returns information about the Notice .
// If user is creator also returns an array of deal for the given notice.
func (controller *NoticeController) GetNoticeInfo(idUser int, noticeId int) (string, []models.Deal, error) {
	notice, err := controller.noticeInteractor.GetNotice(noticeId)
	if err != nil {
		return "", nil, err
	}
	var info string
	var deals []models.Deal
	if notice.ClientId != idUser {
		info, err = controller.noticeInteractor.GetNoticeInfoShort(noticeId)
	} else {
		info, err = controller.noticeInteractor.GetNoticeInfoFull(noticeId)
		if err == nil {
			deals, err = controller.dealInteractor.GetAllNoticeDeals(noticeId)
		}
	}
	return info, deals, err
}
func (controller *NoticeController) GetAllNotices() []models.Notice {
	res := controller.noticeInteractor.GetAllNotices()
	return res
}
func (controller *NoticeController) DeleteNotice(userId int, noticeId int) error {

	notice, err := controller.noticeInteractor.GetNotice(noticeId)
	if err != nil {
		return err
	}

	if notice.ClientId != userId {
		return errors.New("not owner")
	}
	// deals, err := controller.dealInteractor.GetAllNoticeDeals(noticeId)
	// for _, deal
	err = controller.noticeInteractor.Delete(noticeId)
	return err
}

func (controller *NoticeController) AddResponse(userId int, noticeId int) error {
	deal := models.Deal{}
	deal.PerformerId = userId
	deal.NoticeId = noticeId
	deal.Approved = false
	return controller.dealInteractor.Add(deal)
}
func (controller *NoticeController) ApproveDeal(performerId, noticeId int) error {
	return controller.dealInteractor.ApproveDeal(performerId, noticeId)
}
func (controller *NoticeController) DeleteDeal(performerId int, noticeId int) error {
	err := controller.dealInteractor.Delete(performerId, noticeId)
	return err
}
func (controller *NoticeController) SelectWithFilter(category int, typeNotion int, minCost int, maxCost int) ([]models.Notice, error) {
	filter := database.SelectOptions{}
	filter.Fill(typeNotion, category, minCost, maxCost)
	return controller.noticeInteractor.SelectWithFilter(filter)
}
