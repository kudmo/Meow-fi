package services

import (
	"Meow-fi/internal/database"
	"Meow-fi/internal/database/interfaces"
	"Meow-fi/internal/models"
	"Meow-fi/internal/services/usercase/controller"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DealController struct {
	Interactor controller.DealInteractor
}

func NewDealController(sqlHandler interfaces.SqlHandler) *DealController {
	return &DealController{
		Interactor: controller.DealInteractor{
			DealRepository: &database.DealRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *DealController) Create(ctx echo.Context) {
	deal := models.Deal{}
	ctx.Bind(&deal)
	controller.Interactor.Add(deal)
	createdDeals := controller.Interactor.GetAllDeals()
	ctx.JSON(201, createdDeals)
	return
}
func (controller *DealController) UpdateDeal(t models.Deal) {

}
func (controller *DealController) GetDeal(PerformerId string, NoticeId string) models.Deal {
	deal, _ := controller.Interactor.GetDeal(PerformerId, NoticeId)
	return deal
}
func (controller *DealController) GetDealInfo(c echo.Context) error {
	PerformerId := c.Param("performer_id")
	NoticeId := c.Param("notice_id")
	deal, err := controller.Interactor.GetDealInfo(PerformerId, NoticeId)
	if err != nil {
		c.NoContent(http.StatusBadRequest)
		return errors.New("Error")
	}
	c.String(http.StatusOK, deal.Performer.Login+" want to do "+deal.Notice.Containing)
	return nil
}
func (controller *DealController) Delete(PerformerId string, NoticeId string) {
	controller.Interactor.Delete(PerformerId, NoticeId)
}
func (controller *DealController) GetAllDeals() []models.Deal {
	res := controller.Interactor.GetAllDeals()
	return res
}
