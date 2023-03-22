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

	userRoutes := app.Group("/user")
	userRoutes.Post("/create", controllers.UserCreate)
	userRoutes.Get("/all", controllers.UserAll)
	return app
}
