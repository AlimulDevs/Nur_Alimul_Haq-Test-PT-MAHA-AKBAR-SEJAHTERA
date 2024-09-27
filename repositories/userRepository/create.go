package userRepository

import (
	dto "maha-akbar-sejahtera/models/dto"
	"maha-akbar-sejahtera/models/model"

	"github.com/jinzhu/copier"
)

func (ui *UserImplementation) Create(input dto.UserRequest) error {
	var newUser model.User
	err := copier.Copy(&newUser, &input)
	if err != nil {
		return err
	}
	err = ui.db.Model(&model.User{}).Create(&newUser).Error
	if err != nil {
		return err
	}
	return nil
}
