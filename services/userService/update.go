package userService

import (
	dto "maha-akbar-sejahtera/models/dto"
)

func (ui *UserImplementation) Update(id string, input dto.UserRequest) error {
	err := ui.repository.Update(id, input)
	if err != nil {
		return err
	}
	return nil
}
