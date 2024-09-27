package database

import (
	mysql "maha-akbar-sejahtera/database/mysql"
	helpers "maha-akbar-sejahtera/helpers"

	"gorm.io/gorm"
)

func DatabaseConnect() *gorm.DB {
	configDatabase := mysql.ConfigDatabase{
		DB_USERNAME: helpers.GetEnv("DB_USERNAME"),
		DB_PASSWORD: helpers.GetEnv("DB_PASSWORD"),
		DB_ADDRESS:  helpers.GetEnv("DB_ADDRESS"),
		DB_NAME:     helpers.GetEnv("DB_NAME"),
	}

	db := configDatabase.InitDatabase()

	return db
}
