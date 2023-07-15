package interfaces

import "gorm.io/gorm"

type SqlHandler interface {
	Create(object interface{}) error
	Update(object interface{}) error
	Where(object interface{}, conds ...interface{}) (tx *gorm.DB)
	Preload(query string, args ...interface{}) (tx *gorm.DB)
	FindAll(object interface{}) error
	DeleteById(object interface{}, id string) error
	SelectById(object interface{}, id string) error
	Order(value interface{}) (tx *gorm.DB)
	Table(name string) (tx *gorm.DB)
}
