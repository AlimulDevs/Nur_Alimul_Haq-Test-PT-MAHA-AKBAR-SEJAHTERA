package main

import (
	routes "maha-akbar-sejahtera/app/routes"
	"maha-akbar-sejahtera/database"
	"maha-akbar-sejahtera/helpers"
)

func main() {
	db := database.DatabaseConnect()

	err := database.DatabaseMigration(db)

	if err != nil {
		panic(err)
	}

	server := routes.RouteService(db)

	server.Logger.Fatal(server.Start(helpers.GetEnv("APP_PORT")))
}
