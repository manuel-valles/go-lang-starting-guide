package services

import "example.com/contacts_api/models"

type UserService interface {
	CreateUser(user *models.User) error
	GetUser(userName *string) (*models.User, error)
	GetUsers() ([]*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(userName *string) error
}