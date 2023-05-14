package database

import (
	"Meow-fi/internal/config"
	"Meow-fi/internal/database/interfaces"
	"Meow-fi/internal/models"
	"time"

	"gorm.io/gorm"
)

type NoticeRepository struct {
	interfaces.SqlHandler
}

type SelectOptions struct {
	Categories int    `query:"categories"`
	Types      int    `query:"types"`
	MaxCost    int    `query:"max_cost"`
	MinCost    int    `query:"min_cost"`
	OrderBy    string `query:"order_by"`
	PageNumber int    `query:"page"`
}

func (filter *SelectOptions) Fill(types, category, minCost, maxCost, pageNumber int, orderBy string) {
	filter.Categories = category
	filter.Types = types
	filter.MinCost = minCost
	filter.MaxCost = maxCost
	filter.PageNumber = pageNumber
	filter.OrderBy = orderBy
}

func (db *NoticeRepository) Store(notice models.Notice) error {
	return db.Create(&notice)
}

func (db *NoticeRepository) CheckClient(userId, noticeId int) (bool, error) {
	var client_id models.Notice
	res := db.Where("id = ?", noticeId).Select("client_id").Find(&client_id)
	if res.RowsAffected == 0 {
		return false, gorm.ErrRecordNotFound
	}
	return client_id.ClientId == userId, res.Error
}

func (db *NoticeRepository) Select() []models.Notice {
	var notices []models.Notice
	db.FindAll(&notices)
	return notices
}
func (db *NoticeRepository) UpdateNotice(notice models.Notice) error {
	return db.Update(&notice)
}
func (db *NoticeRepository) SelectById(id int) (models.Notice, error) {
	var notice models.Notice
	res := db.Preload("Client").Where("id = ?", id).Find(&notice)
	if res.RowsAffected == 0 {
		return notice, gorm.ErrRecordNotFound
	}
	return notice, res.Error
}
func (db *NoticeRepository) Delete(id int) error {
	var notices []models.Notice
	// return db.Where("id = ?", id).Delete(&notices).Error
	return db.Where("id = ?", id).Unscoped().Delete(&notices).Error
}

func (db *NoticeRepository) SelectWithFilter(filter SelectOptions) ([]models.Notice, error) {
	var notices []models.Notice
	var category models.Category
	var categoties *gorm.DB
	var res *gorm.DB
	res = nil
	if filter.Categories != 0 {
		categoties = db.Where("category_id = ?", filter.Categories).Take(&category)
		if categoties.Error != nil {
			return nil, categoties.Error
		}
		categoties = db.Table("categories").Where("left_key >= ? AND right_key <= ?", category.LeftKey, category.RightKey).Select("category_id")
		res = db.Where("category IN (?)", categoties)
		if res.Error != nil {
			return nil, res.Error
		}
	}
	if filter.Types != 0 {
		if res != nil {
			res = res.Where("type_notice = ?", filter.Types)
		} else {
			res = db.Where("type_notice = ?", filter.Types)
		}
		if res.Error != nil {
			return nil, res.Error
		}
	}
	if filter.MaxCost != 0 {
		if res != nil {
			res = res.Where("cost <= ?", filter.MaxCost)
		} else {
			res = db.Where("cost <= ?", filter.MaxCost)
		}
		if res.Error != nil {
			return nil, res.Error
		}
	}

	if res != nil {
		res = res.Where("cost >= ?", filter.MinCost)
	} else {
		res = db.Where("cost >= ?", filter.MinCost)
	}
	if res.Error != nil {
		return nil, res.Error
	}
	if filter.OrderBy == "" {
		filter.OrderBy = "created_at"
	}
	res = res.Order(filter.OrderBy).Where("time_avaliable >= ?", time.Now()).Offset(config.SizeNotionPage * filter.PageNumber).Limit(config.SizeNotionPage).Find(&notices)
	if res.Error != nil {
		return nil, res.Error
	}
	return notices, nil
}
