package models

type Material struct {
	Id        int    `gorm:"primary_key"`
	CreatorId int    `json:"creator"`
	Name      string `json:"filename`
	Path      string `json:"-" gorm:"NOT NULL"`
	Category  int    `json:"category"`
}
