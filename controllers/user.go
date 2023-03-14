package controllers

import (
	"github.com/arshad-siddiqui/sidequest/models"
	"github.com/gofiber/fiber/v2"
)

func UserCreateController(c *fiber.Ctx) error {
	user := new(models.User)
	err := c.BodyParser(&user)
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.SendString("UserCreateController")
}
