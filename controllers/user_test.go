package controllers

import (
	"net/http"
	"testing"

	"github.com/arshad-siddiqui/sidequest/initialize"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	t.Log("TestUserCreate")

	app := fiber.New()
	app.Post("/user/create", UserCreate)

	req, err := http.NewRequest("POST", "/user/create", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, resp.StatusCode, 400, "Status code should be 400 with no body")
}

func TestUserAll(t *testing.T) {
	t.Log("TestUserAll")

	app := fiber.New()
	app.Get("/user/all", UserAll)

	initialize.LoadTestEnv()
	initialize.ConnectDB()

	req, err := http.NewRequest("GET", "/user/all", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, resp.StatusCode, 200, "Status code should be 200 with no body")
}
