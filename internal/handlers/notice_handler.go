package handlers

import (
	"net/http"

	"Meow-fi/internal/auth"
	controllers "Meow-fi/internal/services"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type NoticeHandler struct {
	Controller controllers.NoticeController
}

func (handler *NoticeHandler) GetNoticeInfo(c echo.Context) {
	id := c.Param("id")
	notice := handler.Controller.GetNotice(id)
	if notice.Client.UserId == c.Get("user").(*jwt.Token).Claims.(*auth.JwtCustomClaims).Id {
		c.JSON(http.StatusOK, "You added notice: \""+notice.Containing+"\" at "+notice.CreatedAt.Format("02-Jan-2006"))
	} else {
		c.JSON(http.StatusOK, "Somebody added notice: \""+notice.Containing+"\"")
	}
}
func (handler *NoticeHandler) GetNoticeInfoGuest(c echo.Context) {
	c.JSON(http.StatusOK, "Not allowed for guests")
}

func (handler *NoticeHandler) UpdateNotice(c echo.Context) {
	id := c.Param("id")
	notice := handler.Controller.GetNotice(id)
	if notice.Client.UserId == c.Get("user").(*jwt.Token).Claims.(*auth.JwtCustomClaims).Id {
		c.Bind(&notice)
		handler.Controller.UpdateNotice(notice)
		c.String(http.StatusOK, "Updated")
	} else {
		c.String(http.StatusOK, "You arn't owner")
	}
}

func (handler *NoticeHandler) DeleteNotice(c echo.Context) {
	id := c.Param("id")
	notice := handler.Controller.GetNotice(id)
	if notice.Client.UserId == c.Get("user").(*jwt.Token).Claims.(*auth.JwtCustomClaims).Id {
		handler.Controller.Delete(id)
		c.String(http.StatusOK, "Deleted")
	} else {
		c.String(http.StatusOK, "You arn't owner")
	}
}
