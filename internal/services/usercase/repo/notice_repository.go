package repo

import "Meow-fi/internal/models"

type NoticeRepository interface {
	Store(models.Notice)
	UpdateNotice(models.Notice)
	Select() []models.Notice
	SelectById(id string) (models.Notice, error)
	Delete(id string)
}
