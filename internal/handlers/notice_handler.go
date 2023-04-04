package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"Meow-fi/internal/auth"
	"Meow-fi/internal/models"
	controllers "Meow-fi/internal/services"

	"github.com/labstack/echo/v4"
)

// Structure for handling requests related to notifications.
// All requests for a particular notice by id use information from the JWT (user id)
type NoticeHandler struct {
	Controller controllers.NoticeController
}

func (handler *NoticeHandler) Create(c echo.Context) error {
	notice := models.Notice{}
	c.Bind(&notice)
	userId := auth.TokenGetUserId(c)
	err := handler.Controller.Create(userId, notice)
	if err != nil {
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.String(http.StatusOK, "created")
}
func (handler *NoticeHandler) GetAllNotices(c echo.Context) error {
	notices := handler.Controller.GetAllNotices()
	return c.JSON(http.StatusOK, notices)
}
func (handler *NoticeHandler) GetNoticeInfo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	userId := auth.TokenGetUserId(c)
	noticeInfo, deals, err := handler.Controller.GetNoticeInfo(userId, id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.JSON(http.StatusOK, echo.Map{
		"info":  noticeInfo,
		"deals": deals,
	})
}
func (handler *NoticeHandler) UpdateNotice(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	notice := models.Notice{}
	c.Bind(&notice)
	userId := auth.TokenGetUserId(c)
	err = handler.Controller.UpdateNotice(userId, id, notice)
	if err != nil {
		if errors.Is(err, errors.New("not owner")) {
			return c.String(http.StatusForbidden, "not owner")
		}
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.String(http.StatusOK, "updated")
}
func (handler *NoticeHandler) DeleteNotice(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	userId := auth.TokenGetUserId(c)
	err = handler.Controller.Delete(userId, id)
	if err != nil {
		if errors.Is(err, errors.New("not owner")) {
			return c.String(http.StatusForbidden, "not owner")
		}
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.String(http.StatusOK, "updated")
}

// Add unapproved deal for notice by id
func (handler *NoticeHandler) AddResponse(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	userId := auth.TokenGetUserId(c)
	err = handler.Controller.AddResponse(userId, id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.String(http.StatusOK, "added")
}
