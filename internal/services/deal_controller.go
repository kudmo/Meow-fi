package services

import (
	"Meow-fi/internal/database"
	"Meow-fi/internal/database/interfaces"
	"Meow-fi/internal/models"
	"Meow-fi/internal/services/usercase/controller"

	"github.com/labstack/echo"
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
	deal := controller.Interactor.GetDeal(PerformerId, NoticeId)
	return deal
}
func (controller *DealController) GetDealInfo(PerformerId string, NoticeId string) string {
	deal := controller.Interactor.GetDealInfo(PerformerId, NoticeId)
	str := ""
	str += deal.Client.FIO + " created Deal: " + deal.Containing
	return str
}
func (controller *DealController) Delete(PerformerId string, NoticeId string) {
	controller.Interactor.Delete(PerformerId, NoticeId)
}
func (controller *DealController) GetAllDeals() []models.Deal {
	res := controller.Interactor.GetAllDeals()
	return res
}
