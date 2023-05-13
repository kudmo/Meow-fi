package services

import (
	"Meow-fi/internal/auth"
	"Meow-fi/internal/database"
	"Meow-fi/internal/database/interfaces"
	"Meow-fi/internal/models"
	"Meow-fi/internal/services/usercase/repo"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

type MaterialController struct {
	MaterialRepository repo.MaterialRepository
}

func NewMaterialController(sqlHandler interfaces.SqlHandler) *MaterialController {
	return &MaterialController{
		MaterialRepository: &database.MaterialRepository{
			SqlHandler: sqlHandler,
		},
	}
}

func (controller *MaterialController) Upload(c echo.Context) error {
	userId := auth.TokenGetUserId(c)
	// userId := 1
	category, err := strconv.Atoi(c.FormValue("category"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request: "+err.Error())
	}
	filename := c.FormValue("filename")
	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request: "+err.Error())
	}

	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "something goes wrong: "+err.Error())
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("materials/" + file.Filename)
	if err != nil {
		return c.String(http.StatusInternalServerError, "something goes wrong: "+err.Error())
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return c.String(http.StatusInternalServerError, "something goes wrong: "+err.Error())
	}

	material := models.Material{
		Path:      "materials/" + file.Filename,
		CreatorId: userId,
		Category:  category,
		Name:      filename}

	err = controller.MaterialRepository.Store(material)
	if err != nil {
		return c.String(http.StatusInternalServerError, "something goes wrong: "+err.Error())
	}

	return c.String(http.StatusCreated, "uploaded")
}
func (controller *MaterialController) Download(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	material, err := controller.MaterialRepository.SelectById(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, "something goes wrong: "+err.Error())
	}
	return c.Inline(material.Path, material.Name)
}
func (controller *MaterialController) SelectWithFilter(c echo.Context) error {
	category, err := strconv.Atoi(c.FormValue("category"))
	if err != nil {
		category = 0
	}
	filter := database.SelectOptions{}
	filter.Fill(0, category, 0, 0)
	res, err := controller.MaterialRepository.SelectWithFilter(filter)
	if err != nil {
		log.Println(err.Error())
		return c.String(http.StatusInternalServerError, "something goes wrong")
	}
	return c.JSON(http.StatusOK, res)
}
