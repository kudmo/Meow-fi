package models

type User struct {
	UserId   int    `json:"-" gorm:"primary_key"`
	Login    string `json:"login" gorm:"unique"`
	Password string `json:"-"`
	Salt     string `json:"-"`
}
