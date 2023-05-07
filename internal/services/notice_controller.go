package services

import (
	"Meow-fi/internal/database"
	"Meow-fi/internal/database/interfaces"
	"Meow-fi/internal/models"
	"Meow-fi/internal/services/usercase/controller"
	"errors"
	"log"
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

	var infoChan chan string = make(chan string, 1)
	var dealsChan chan []models.Deal = make(chan []models.Deal, 1)
	var errInfoChan chan error = make(chan error, 2)
	var errDealsChan chan error = make(chan error, 2)

	if notice.ClientId != idUser {
		go func() {
			info, err := controller.noticeInteractor.GetNoticeInfoShort(noticeId)
			infoChan <- info
			errInfoChan <- err
		}()
		errDealsChan <- nil
		dealsChan <- nil
	} else {
		go func() {
			if len(errDealsChan) != 0 {
				if err := <-errDealsChan; err != nil {
					errDealsChan <- err
					errInfoChan <- nil
					return
				} else {
					errDealsChan <- err
				}
			}
			info, err := controller.noticeInteractor.GetNoticeInfoFull(noticeId)
			infoChan <- info
			errInfoChan <- err
		}()
		go func() {
			if len(errInfoChan) != 0 {
				err := <-errInfoChan
				if err != nil {
					errDealsChan <- nil
					errInfoChan <- err

					return
				} else {
					errInfoChan <- err

				}
			}
			deals, err := controller.dealInteractor.GetAllNoticeDeals(noticeId)
			dealsChan <- deals
			errDealsChan <- err
		}()
	}
	if err := <-errInfoChan; err != nil {
		log.Println("Error: in getting information")
		return "", nil, err
	}
	if err := <-errDealsChan; err != nil {
		log.Println("Error: in getting deals")
		return "", nil, err
	}
	return <-infoChan, <-dealsChan, nil
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
func (controller *NoticeController) SelectWithFilter(category int, typeNotion int) ([]models.Notice, error) {
	filter := database.SelectOptions{}
	filter.Fill(typeNotion, category)
	return controller.noticeInteractor.SelectWithFilter(filter)
}
