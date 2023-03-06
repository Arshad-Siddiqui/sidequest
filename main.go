package main

import (
	"github.com/arshad-siddiqui/sidequest/initialize"
	"github.com/gofiber/fiber/v2"
)

func init() {
	initialize.LoadEnv()
}
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Listen(":3000")

}
