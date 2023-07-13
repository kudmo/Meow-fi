package models

type Category struct {
	CategoryId     int `gorm:"primary_key"`
	ParentCategory int
	Level          int
	LeftKey        int
	RightKey       int
	Name           string
}
