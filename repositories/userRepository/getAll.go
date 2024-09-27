package userRepository

import (
	dto "maha-akbar-sejahtera/models/dto"
	"maha-akbar-sejahtera/models/model"

	"github.com/jinzhu/copier"
)

func (ui *UserImplementation) GetAll() ([]dto.UserResponse, error) {
	var response []dto.UserResponse
	var userModel []model.User
	err := ui.db.Find(&userModel).Error
	if err != nil {
		return []dto.UserResponse{}, err
	}

	err = copier.Copy(&response, &userModel)
	if err != nil {
		return []dto.UserResponse{}, err
	}

	return response, nil
}
