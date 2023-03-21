package controllers

import (
	"github.com/arshad-siddiqui/sidequest/initialize"
	"github.com/arshad-siddiqui/sidequest/models"
	"github.com/gofiber/fiber/v2"
)

func UserCreate(c *fiber.Ctx) error {
	user := new(models.User)
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	initialize.DB.Create(&user)
	return c.SendString("UserCreateController")
}

func UserAll(c *fiber.Ctx) error {
	var users []models.User
	initialize.DB.Find(&users)
	return c.JSON(users)
}

func UserDelete(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
