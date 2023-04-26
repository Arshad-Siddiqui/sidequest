package controllers

import (
	"encoding/json"

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
	var data map[string]interface{}
	err := json.Unmarshal(c.Body(), &data)
	if err != nil {
		c.Status(400).SendString(err.Error())
	}

	email := data["email"].(string)
	// email, ok := emailVal.(string)
	// if !ok {
	// 	// handle invalid email value
	// }

	// finding the user with the email
	var user models.User
	initialize.DB.Where("email = ?", email).First(&user)
	if user.Email == "" {
		return c.Status(404).SendString("User not found")
	}

	// returning the user
	return c.JSON(user)
}
