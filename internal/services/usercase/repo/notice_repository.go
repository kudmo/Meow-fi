package repo

import "Meow-fi/internal/models"

type NoticeRepository interface {
	Store(models.Notice) error
	UpdateNotice(models.Notice) error
	Select() []models.Notice
	SelectById(id string) (models.Notice, error)
	Delete(id string) error
}
