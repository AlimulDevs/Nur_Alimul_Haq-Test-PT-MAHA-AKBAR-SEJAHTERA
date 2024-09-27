package userService

import (
	dto "maha-akbar-sejahtera/models/dto"
)

func (ui *UserImplementation) GetAll() ([]dto.UserResponse, error) {
	response, err := ui.repository.GetAll()
	if err != nil {
		return []dto.UserResponse{}, err
	}
	return response, nil
}
