package database

import (
	"Meow-fi/internal/config"
	"Meow-fi/internal/database/interfaces"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SqlHandler struct {
	db *gorm.DB
}

func NewSqlHandler() interfaces.SqlHandler {
	dsn := config.DatabaseUrl
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.db = db
	return sqlHandler
}
func (handler *SqlHandler) Create(obj interface{}) error {
	err := handler.db.Create(obj).Error
	return err
}
func (handler *SqlHandler) Update(obj interface{}) error {
	err := handler.db.Save(obj).Error
	return err
}
func (handler *SqlHandler) FindAll(obj interface{}) error {
	err := handler.db.Find(obj).Error
	return err
}
func (handler *SqlHandler) DeleteById(obj interface{}, id string) error {
	err := handler.db.Delete(obj, id).Error
	return err

}
func (handler *SqlHandler) SelectById(obj interface{}, id string) error {
	err := handler.db.Select(obj, id).Error
	return err
}
func (handler *SqlHandler) Where(object interface{}, args ...interface{}) (tx *gorm.DB) {
	return handler.db.Where(object, args)
}
func (handler *SqlHandler) Preload(query string, args ...interface{}) (tx *gorm.DB) {
	return handler.db.Preload(query, args)
}
