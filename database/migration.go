package database

import (
	model "maha-akbar-sejahtera/models/model"

	"gorm.io/gorm"
)

func DatabaseMigration(database *gorm.DB) error {
	err := database.AutoMigrate(
		&model.User{},
	)

	if err != nil {
		return err
	}

	return nil
}
