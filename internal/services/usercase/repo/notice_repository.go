package repo

import "Meow-fi/internal/models"

type NoticeRepository interface {
	Store(models.Notice) error
	UpdateNotice(models.Notice) error
	Select() []models.Notice
	SelectById(id int) (models.Notice, error)
	CheckClient(userId, noticeId int) (bool, error)
	Delete(id int) error
}
