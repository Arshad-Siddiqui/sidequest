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

	addUserRoutes(app)
	return app
}

func addUserRoutes(app *fiber.App) {
	userRoutes := app.Group("/user")
	userRoutes.Post("/create", controllers.UserCreate)
	userRoutes.Post("/find", controllers.UserFind)
	userRoutes.Get("/all", controllers.UserAll)
}
