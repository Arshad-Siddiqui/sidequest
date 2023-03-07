package main

import (
	"log"
	"os"

	"github.com/arshad-siddiqui/sidequest/initialize"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initialize.LoadEnv()
	initialize.ConnectDB()
}
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Listen(":3000")
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
