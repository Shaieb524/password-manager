package main

import (
	"fmt"
	apController "password-manager/src/controllers/accountPassword"
	apRepo "password-manager/src/models/database/accountPassword"
	"password-manager/src/providers/database"
	"password-manager/src/routes"
	apService "password-manager/src/services/accountPassword"
)

func main() {
	db := database.NewDatabaseContext()
	repo := apRepo.ProvideModuleforDI(db)
	service := apService.ProvideModuleforDI(repo)
	apCC := apController.ProvideModuleforDI(service)

	fmt.Println("Start bitch")
	routes.LoadEnv()
	routes.LoadDatabase()
	routes.SetupRoutesAndRun(apCC)
}
