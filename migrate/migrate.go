package main

import (
	"fmt"

	"github.com/arshad-siddiqui/sidequest/initialize"
	"github.com/arshad-siddiqui/sidequest/models"
)

func init() {
	initialize.LoadEnv()
	initialize.ConnectDB()
}

func main() {
	initialize.DB.AutoMigrate(&models.User{})
	fmt.Println("Migrated")
}
