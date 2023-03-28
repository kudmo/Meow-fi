package models

type User struct {
	UserId   int    `json:"user_id" gorm:"primary_key"`
	Login    string `json:"login" gorm:"unique"`
	Password string
	Salt     string
}
