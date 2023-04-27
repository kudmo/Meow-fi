package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"Meow-fi/internal/auth"
	"Meow-fi/internal/models"
	controllers "Meow-fi/internal/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Structure for handling requests related to notifications.
// All requests for a particular notice by id use information from the JWT (user id)
type NoticeHandler struct {
	Controller controllers.NoticeController
}

// Create new notice
// The client id is the user id from the token
// Other parameters are obtained from JSON
func (handler *NoticeHandler) CreateNotice(c echo.Context) error {
	notice := models.Notice{}
	c.Bind(&notice)
	userId := auth.TokenGetUserId(c)
	err := handler.Controller.Create(userId, notice)
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.String(http.StatusCreated, "created")
}

// Returns all notices
func (handler *NoticeHandler) GetAllNotices(c echo.Context) error {
	notices := handler.Controller.GetAllNotices()
	return c.JSON(http.StatusOK, notices)
}

// Returns information about the notice
// For different users - a different amount of information
// Also returns an array of deals for the creator
func (handler *NoticeHandler) GetNoticeInfo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	userId := auth.TokenGetUserId(c)
	noticeInfo, deals, err := handler.Controller.GetNoticeInfo(userId, id)
	if err == gorm.ErrRecordNotFound {
		return c.String(http.StatusNotFound, "no notice")
	}
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"info":  noticeInfo,
		"deals": deals,
	})
}

// Update notice
// Only the creator can update the notice
func (handler *NoticeHandler) UpdateNotice(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	notice := models.Notice{}
	c.Bind(&notice)
	userId := auth.TokenGetUserId(c)
	err = handler.Controller.UpdateNotice(userId, id, notice)
	if errors.Is(err, errors.New("not owner")) {
		return c.String(http.StatusForbidden, "not owner")
	}
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.String(http.StatusOK, "updated")
}

// Deletee notice
// Only the creator can delete the notice
func (handler *NoticeHandler) DeleteNotice(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	userId := auth.TokenGetUserId(c)
	err = handler.Controller.DeleteNotice(userId, id)
	if err == gorm.ErrRecordNotFound {
		return c.String(http.StatusNotFound, "no notice")
	}
	if errors.Is(err, errors.New("not owner")) {
		return c.String(http.StatusForbidden, "not owner")
	}
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.String(http.StatusOK, "deleted")
}

// Add unapproved deal for notice by id
// The performer id is the user id from the token
func (handler *NoticeHandler) AddResponse(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	userId := auth.TokenGetUserId(c)
	err = handler.Controller.AddResponse(userId, id)
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.String(http.StatusOK, "added")
}

// Approve deal
// Only the creator of notice can approve the deal
func (handler *NoticeHandler) ApproveResponse(c echo.Context) error {
	notice_id, err := strconv.Atoi(c.Param("notice_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	performer_id, err := strconv.Atoi(c.Param("performer_id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	userId := auth.TokenGetUserId(c)
	flag, err := handler.Controller.CheckClient(userId, notice_id)
	if err == gorm.ErrRecordNotFound {
		return c.String(http.StatusNotFound, "no notice")
	}
	if !flag {
		return c.String(http.StatusForbidden, "not allowed")
	}

	err = handler.Controller.ApproveDeal(performer_id, notice_id)
	if err == gorm.ErrRecordNotFound {
		return c.String(http.StatusNotFound, "no deal")
	}
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.String(http.StatusOK, "approved")
}

// Delete deal
// By param (notice_id) in url and user id
// The user can only delete their own deal
func (handler *NoticeHandler) DeleteDeal(c echo.Context) error {
	notice_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	userId := auth.TokenGetUserId(c)

	err = handler.Controller.DeleteDeal(userId, notice_id)
	if err == gorm.ErrRecordNotFound {
		return c.String(http.StatusNotFound, "no deal")
	}
	if err == nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.String(http.StatusOK, "deleted")
}

// Returns all notions by GET-params (default value - 0)
func (handler *NoticeHandler) SelectWithFilter(c echo.Context) error {
	category, err := strconv.Atoi(c.QueryParam("category"))
	if err != nil {
		category = 0
	}
	typeNotion, err := strconv.Atoi(c.QueryParam("type"))
	if err != nil {
		typeNotion = 0
	}
	res, err := handler.Controller.SelectWithFilter(category, typeNotion)
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.JSON(http.StatusOK, res)
}
