package userRepository

import (
	"maha-akbar-sejahtera/models/model"
)

func (ui *UserImplementation) Delete(id string) error {
	err := ui.db.Where("id = ?", id).Delete(&model.User{}).Error

	return err
}
