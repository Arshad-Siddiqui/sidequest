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

	if user.Email == "" || user.Password == "" {
		return c.Status(400).SendString("Email and Password are required")
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

func UserFind(c *fiber.Ctx) error {
	// getting the email from the post request json
	email := c.Params("email")

	// finding the user with the email
	var user models.User
	initialize.DB.Where("email = ?", email).First(&user)

	// returning the user
	return c.JSON(user)
}
