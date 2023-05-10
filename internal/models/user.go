package models

type User struct {
	Id               int    `json:"-" gorm:"primary_key"`
	Email            string `json:"-" gorm:"unique"`
	Login            string `json:"login" gorm:"unique"`
	Password         string `json:"-"`
	Salt             string `json:"-"`
	CurrentRefreshId string `json:"-"`
}
