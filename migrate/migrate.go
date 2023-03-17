package main

import (
	"fmt"

	"github.com/arshad-siddiqui/sidequest/initialize"
	"github.com/arshad-siddiqui/sidequest/models"
)

func main() {
	setDB(".env.test")
	fmt.Println("Migrated Test DB")

	setDB(".env")
	fmt.Println("Migrated Local DB")
}

func setDB(environment string) {
	initialize.LoadEnv(environment)
	initialize.ConnectDB()
	initialize.DB.AutoMigrate(&models.User{})
}
