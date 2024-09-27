package userService

import (
	dto "maha-akbar-sejahtera/models/dto"
)

func (ui *UserImplementation) Create(input dto.UserRequest) error {
	err := ui.repository.Create(input)
	if err != nil {
		return err

	}

	return nil
}
