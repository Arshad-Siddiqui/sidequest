package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func TestController(c *fiber.Ctx) error {
	type Test struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	var test = Test{
		Name:     "Arshad",
		Password: "123456",
	}

	json, err := json.Marshal(test)
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.SendString(string(json))
}
