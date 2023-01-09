package main

import (
	"fmt"
	"time"

	authController "password-manager/src/controllers/authentication"
	userRepo "password-manager/src/models/database/user"
	authService "password-manager/src/services/authentcation"

	apController "password-manager/src/controllers/accountPassword"
	apRepo "password-manager/src/models/database/accountPassword"
	apService "password-manager/src/services/accountPassword"

	LocalCache "password-manager/src/providers/appCache"
	"password-manager/src/providers/database"
	"password-manager/src/routes"
	"password-manager/src/utils/env"
	// web "password-manager/src/web"
)

func main() {
	fmt.Println("Start bitch")

	globalEnv := env.NewEnv()
	db := database.NewDatabaseContext()
	lCache := LocalCache.NewLocalCache(20 * time.Minute)

	authR := userRepo.NewAuthenticationRepoModule(db)
	authS := authService.NewAuthenticationServiceModule(authR)
	authC := authController.NewAuthenticationControllerModule(authS)

	apR := apRepo.NewAccPasswordRepoModule(db)
	apS := apService.NewAccPasswordServiceModule(apR, lCache)
	apCC := apController.NewAccPasswordControllerModule(apS)

	// api := web.InitAccountPasswordAPI()
	routes.SetupRoutesAndRun(apCC, authC, globalEnv)
}
