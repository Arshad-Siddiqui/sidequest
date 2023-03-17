package main

import (
	"fmt"

	"github.com/arshad-siddiqui/sidequest/initialize"
	"github.com/arshad-siddiqui/sidequest/models"
)

func main() {
	initialize.LoadEnv(".env.test")
	initialize.ConnectDB()

	initialize.DB.AutoMigrate(&models.User{})
	fmt.Println("Migrated Test DB")

	initialize.LoadEnv(".env")
	initialize.ConnectDB()

	initialize.DB.AutoMigrate(&models.User{})
	fmt.Println("Migrated Local DB")
}
