package repository

import "crud/internal/domain/models"

type Service interface {
	Init() error
	CreateUser(user models.User) error
	ReadUser(id string) (*models.User, error)
	UpdateUser(user models.User) (*models.User, error)
	DeleteUser(id string) (*models.User, error)
	Stop() error
}
