package userRepository

import (
	dto "maha-akbar-sejahtera/models/dto"
	"maha-akbar-sejahtera/models/model"

	"github.com/jinzhu/copier"
)

func (ui *UserImplementation) Update(id string, input dto.UserRequest) error {
	var updateUser model.User

	err := copier.Copy(&updateUser, &input)
	if err != nil {
		return err
	}
	err = ui.db.Model(&model.User{}).Where("id = ?", id).Updates(&updateUser).Error
	if err != nil {
		return err
	}

	return nil
}
