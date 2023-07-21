package models

type User struct {
	Id          int    `json:"-" gorm:"primary_key;autoIncrement:false;not null;"`
	Login       string `json:"login"`
	Description string `json:"description"`
}
