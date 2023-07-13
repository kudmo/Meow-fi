package repo

import (
	"Meow-fi_app-back/internal/database"
	"Meow-fi_app-back/internal/models"
)

type NoticeRepository interface {
	Store(models.Notice) error
	UpdateNotice(models.Notice) error
	Select() []models.Notice
	SelectById(id int) (models.Notice, error)
	CheckClient(userId, noticeId int) (bool, error)
	Delete(id int) error
	SelectWithFilter(filter database.SelectOptions) ([]models.Notice, error)
}
