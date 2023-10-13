package database

import "github.com/giovannymassuia/learning-go/07-api/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
