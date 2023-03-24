package database

import (
	"Meow-fi/internal/database/interfaces"
	"Meow-fi/internal/models"
)

type UserRepository struct {
	interfaces.SqlHandler
}

func (db *UserRepository) Store(u models.User) {
	db.Create(&u)
}

func (db *UserRepository) Select() []models.User {
	var user []models.User
	db.FindAll(&user)
	return user
}

func (db *UserRepository) Delete(id string) {
	var user []models.User
	db.DeleteById(&user, id)
}
