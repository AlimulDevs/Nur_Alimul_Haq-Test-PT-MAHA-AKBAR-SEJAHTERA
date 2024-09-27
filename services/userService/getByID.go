package userService

import (
	dto "maha-akbar-sejahtera/models/dto"
)

func (ui *UserImplementation) GetByID(id string) (dto.UserResponse, error) {
	response, err := ui.repository.GetByID(id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return response, nil
}
