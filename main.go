package main

import (
	"fmt"
	apController "password-manager/src/controllers/accountPassword"
	apRepo "password-manager/src/models/database/accountPassword"
	"password-manager/src/providers/database"
	"password-manager/src/routes"

	// web "password-manager/src/web"
	apService "password-manager/src/services/accountPassword"
)

func main() {

	fmt.Println("Start bitch")
	routes.LoadEnv()
	routes.LoadDatabase()
	db := database.NewDatabaseContext()
	repo := apRepo.NewAccPasswordRepoModule(db)
	service := apService.NewAccPasswordServiceModule(repo)
	apCC := apController.NewAccPasswordControllerModule(service)

	// api := web.InitAccountPasswordAPI()
	routes.SetupRoutesAndRun(apCC)
}
