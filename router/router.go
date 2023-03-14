package router

import (
	"github.com/arshad-siddiqui/sidequest/controllers"
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/test", controllers.TestController)
	app.Post("/user/create", controllers.UserCreateController)

	return app
}
