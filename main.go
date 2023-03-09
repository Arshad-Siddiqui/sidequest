package main

import (
	"log"
	"os"

	"github.com/arshad-siddiqui/sidequest/initialize"
	"github.com/arshad-siddiqui/sidequest/router"
)

func init() {
	initialize.LoadEnv()
	initialize.ConnectDB()
}
func main() {
	app := router.New()
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
