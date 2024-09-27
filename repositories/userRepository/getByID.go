package userRepository

import (
	dto "maha-akbar-sejahtera/models/dto"
	"maha-akbar-sejahtera/models/model"

	"github.com/jinzhu/copier"
)

func (ui *UserImplementation) GetByID(id string) (dto.UserResponse, error) {
	var response dto.UserResponse
	var userModel model.User
	err := ui.db.Where("id = ?", id).First(&userModel).Error
	if err != nil {
		return dto.UserResponse{}, err
	}
	err = copier.Copy(&response, &userModel)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return response, nil
}
