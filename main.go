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

	"password-manager/src/utils/logger"
)

func main() {
	fmt.Println("Start bitch")

	logger := logger.NewLogger()
	globalEnv := env.NewEnv()
	db := database.NewDatabaseContext()
	lCache := LocalCache.NewLocalCache(1 * time.Minute)

	authR := userRepo.NewAuthenticationRepoModule(logger, db)
	authS := authService.NewAuthenticationServiceModule(logger, authR)
	authC := authController.NewAuthenticationControllerModule(logger, authS)

	apR := apRepo.NewAccPasswordRepoModule(logger, db)
	apS := apService.NewAccPasswordServiceModule(logger, apR, lCache)
	apCC := apController.NewAccPasswordControllerModule(logger, apS)

	// api := web.InitAccountPasswordAPI()
	routes.SetupRoutesAndRun(apCC, authC, globalEnv)
}
