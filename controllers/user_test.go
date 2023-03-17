package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/arshad-siddiqui/sidequest/initialize"
	"github.com/arshad-siddiqui/sidequest/reset"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	t.Log("TestUserCreate")

	initialize.LoadEnv("../.env.test")
	initialize.ConnectDB()
	reset.ResetDB()

	url := "/user/create"
	app := fiber.New()
	app.Post(url, UserCreate)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, resp.StatusCode, 400, "Status code should be 400 with no body")

	email, password := "testemail", "testpassword"
	values := map[string]string{"email": email, "password": password}

	jsonValue, err := json.Marshal(values)
	if err != nil {
		t.Fatal(err)
	}

	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err = app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, resp.StatusCode, 200, "Status code should be 200 with body")
}

func TestUserAll(t *testing.T) {
	t.Log("TestUserAll")

	initialize.LoadEnv("../.env.test")
	initialize.ConnectDB()
	reset.ResetDB()

	url := "/user/all"
	app := fiber.New()
	app.Get(url, UserAll)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, resp.StatusCode, 200, "Status code should be 200 with no body")
}
