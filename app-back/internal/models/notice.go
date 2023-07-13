package models

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Notice struct {
	Id            int       `json:"id" gorm:"primary_key"`
	TypeNotice    int       `json:"type" gorm:"not null"`
	ClientId      int       `json:"-" gorm:"not null"`
	Containing    string    `json:"containing" gorm:"not null"`
	Category      int       `json:"category"`
	Cost          int       `json:"cost"`
	Client        User      `json:"-" gorm:"not null"`
	TimeAvaliable time.Time `json:"time_avaliable"`
	CreatedAt     time.Time `json:"-"`
}

func (n *Notice) AfterCreate(tx *gorm.DB) (err error) {
	log.Printf("Created notice %d\n", n.Id)
	return
}

func (n *Notice) BeforeDelete(tx *gorm.DB) (err error) {
	log.Printf("Trying delete notice %d\n", n.Id)
	return
}

func (n *Notice) AfterDelete(tx *gorm.DB) (err error) {
	log.Printf("Deleted notice %d\n", n.Id)
	return
}
