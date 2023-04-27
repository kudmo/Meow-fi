package models

import (
	"log"

	"gorm.io/gorm"
)

type Deal struct {
	PerformerId int    `json:"performer_id" gorm:"primary_key;autoIncrement:false;not null"`
	NoticeId    int    `json:"notice_id" gorm:"primary_key;autoIncrement:false;not null"`
	Approved    bool   `json:"approved"`
	Performer   User   `json:"performer" gorm:"foreignKey:PerformerId;not null;constraint:OnDelete:CASCADE;"`
	Notice      Notice `json:"notice" gorm:"foreignKey:NoticeId;not null;association_foreignkey:ID;constraint:OnDelete:CASCADE;"`
}

func (d *Deal) AfterDelete(tx *gorm.DB) (err error) {
	log.Printf("Deleted deal user:%d for notice %d\n", d.PerformerId, d.NoticeId)
	return
}
