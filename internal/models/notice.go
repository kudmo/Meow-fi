package models

import "time"

type Notice struct {
	Id            int       `json:"id" gorm:"primary_key"`
	TypeNotice    int       `json:"type" gorm:"not null"`
	ClientId      int       `json:"-" gorm:"not null"`
	Containing    string    `json:"containing" gorm:"not null"`
	Category      int       `json:"category"`
	Cost          int       `json:"cost"`
	Client        User      `json:"client" gorm:"not null"`
	TimeAvaliable time.Time `json:"time_avaliable"`
	CreatedAt     time.Time `json:"-"`
}
