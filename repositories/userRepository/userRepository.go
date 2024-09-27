package userRepository

import (
	dto "maha-akbar-sejahtera/models/dto"

	"gorm.io/gorm"
)

type UserImplementation struct {
	db *gorm.DB
}

type UserRepository interface {
	GetAll() ([]dto.UserResponse, error)
	GetByID(id string) (dto.UserResponse, error)
	Create(input dto.UserRequest) error
	Update(id string, input dto.UserRequest) error
	Delete(id string) error
}

func NewUserRepositoy(db *gorm.DB) UserRepository {
	return &UserImplementation{
		db: db,
	}
}
