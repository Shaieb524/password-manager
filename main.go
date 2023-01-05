package main

import (
	"fmt"
	// apController "password-manager/src/controllers/accountPassword"
	// apRepo "password-manager/src/models/database/accountPassword"
	// "password-manager/src/providers/database"
	"password-manager/src/routes"
	wire "password-manager/src/wire"
	// apService "password-manager/src/services/accountPassword"
)

func main() {
	// db := database.NewDatabaseContext()
	// repo := apRepo.NewAccPasswordRepoModule(db)
	// service := apService.NewAccPasswordServiceModule(repo)
	// apCC := apController.NewAccPasswordControllerModule(service)
	fmt.Println("Start bitch")
	routes.LoadEnv()
	routes.LoadDatabase()
	api := wire.InitAccountPasswordAPI()
	routes.SetupRoutesAndRun(api)
}
