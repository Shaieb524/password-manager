package main

import (
	"fmt"
	"password-manager/src/routes"
)

func main() {
	fmt.Println("Start bitch")
	routes.LoadEnv()
	routes.LoadDatabase()
	routes.SetupRoutesAndRun()
}
