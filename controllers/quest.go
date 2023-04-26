package controllers

import (
	"github.com/arshad-siddiqui/sidequest/initialize"
	"github.com/arshad-siddiqui/sidequest/models"
	"github.com/gofiber/fiber/v2"
)

func QuestCreate(c *fiber.Ctx) error {
	quest := new(models.Quest)
	err := c.BodyParser(&quest)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}

	if quest.Title == "" || quest.Description == "" {
		return c.Status(400).SendString("Title and Description are required")
	}

	initialize.DB.Create(&quest)
	return c.SendString("QuestCreateController")
}
