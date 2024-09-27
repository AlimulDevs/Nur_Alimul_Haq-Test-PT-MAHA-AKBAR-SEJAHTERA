package userService

import (
	"maha-akbar-sejahtera/models/dto"
	userRepository "maha-akbar-sejahtera/repositories/userRepository"
)

type UserService interface {
	GetAll() ([]dto.UserResponse, error)
	GetByID(id string) (dto.UserResponse, error)
	Create(input dto.UserRequest) error
	Update(id string, input dto.UserRequest) error
	Delete(id string) error
}

type UserImplementation struct {
	repository userRepository.UserRepository
}

func NewUserService(repository userRepository.UserRepository) UserService {
	return &UserImplementation{
		repository: repository,
	}
}
