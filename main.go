package main

import (
	"fmt"
	"password-manager/src/controllers/accountPassword"
	"password-manager/src/routes"
)

func main() {
	var apc accountPassword.AccountPasswordController
	fmt.Println("Start bitch")
	routes.LoadEnv()
	routes.LoadDatabase()
	routes.SetupRoutesAndRun(apc)
}
