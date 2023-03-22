package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/arshad-siddiqui/sidequest/controllers"
	. "github.com/arshad-siddiqui/sidequest/initialize"
	"github.com/arshad-siddiqui/sidequest/reset"
	"github.com/gofiber/fiber/v2"
)

func addUser(email, password string) {
	initDB()

	url := "/user/create"
	app := fiber.New()
	app.Post(url, controllers.UserCreate)

	values := map[string]string{"email": email, "password": password}

	jsonValue, _ := json.Marshal(values)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))

	req.Header.Set("Content-Type", "application/json")

	app.Test(req)
}

func initAndResetDB() {
	initDB()
	reset.ResetDB()
}

func initDB() {
	LoadEnv("../.env.test")
	ConnectDB()
}
