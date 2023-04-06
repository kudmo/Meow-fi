package models

type Deal struct {
	PerformerId int    `json:"performer_id" gorm:"primary_key;autoIncrement:false;not null"`
	NoticeId    int    `json:"notice_id" gorm:"primary_key;autoIncrement:false;not null"`
	Approved    bool   `json:"approved"`
	Performer   User   `json:"performer" gorm:"foreignKey:PerformerId;not null"`
	Notice      Notice `json:"notice" gorm:"foreignKey:NoticeId;not null"`
}
