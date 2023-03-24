package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/arshad-siddiqui/sidequest/controllers"
	"github.com/arshad-siddiqui/sidequest/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	t.Log("TestUserCreate")

	initAndResetDB()

	url := "/user/create"
	app := fiber.New()
	app.Post(url, controllers.UserCreate)

	tests := []struct {
		email       string
		password    string
		statusCode  int
		description string
	}{
		{
			email:       "testemail",
			password:    "testpassword",
			statusCode:  200,
			description: "Should create user with email and password",
		},
		{
			email:       "testemail2",
			password:    "testpassword2",
			statusCode:  200,
			description: "Should create user with email and password",
		},
		{
			email:       "",
			password:    "",
			statusCode:  400,
			description: "Should not create user with empty email and password",
		},
	}

	for _, test := range tests {
		values := map[string]string{"email": test.email, "password": test.password}

		jsonValue, err := json.Marshal(values)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, test.statusCode, resp.StatusCode, test.description)
	}
}

func TestUserAll(t *testing.T) {
	t.Log("TestUserAll")

	initAndResetDB()

	url := "/user/all"
	app := fiber.New()
	app.Get(url, controllers.UserAll)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, resp.StatusCode, "Status code should be 200 with no body")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "[]", string(body), "Body should be empty array")

	addUser("testemail", "testpassword")
	resp, err = app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, resp.StatusCode, "Status code should be 200 with no body")
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var users []models.User
	err = json.Unmarshal(body, &users)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(users), "Body should be array with one user")
	assert.Equal(t, "testemail", users[0].Email, "User email should be testemail")
	assert.Equal(t, "testpassword", users[0].Password, "User email should be testpassword")

	addUser("testemail2", "testpassword2")

	resp, err = app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(body, &users)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(users), "Body should be array with two users")
	assert.Equal(t, "testemail", users[0].Email, "User email should be testemail")
	assert.Equal(t, "testemail2", users[1].Email, "User2 email should be testemail2")
}

func TestUserDelete(t *testing.T) {
	t.Log("TestUserDelete")

	initAndResetDB()

	url := "/user/delete"
	app := fiber.New()
	app.Delete(url, controllers.UserDelete)

	addUser("testemail", "testpassword")

	values := map[string]string{"email": "testemail", "password": "testpassword"}
	jsonValue, err := json.Marshal(values)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, resp.StatusCode, "Status code should be 200 with no body")
}

func TestUserFind(t *testing.T) {
	tests := []struct {
		email       string
		password    string
		statusCode  int
		description string
	}{
		{
			email:       "testemail1",
			password:    "testpassword1",
			statusCode:  200,
			description: "Should find user with email",
		},
		{
			email:       "testemail3",
			password:    "testpassword3",
			statusCode:  200,
			description: "Should find user with email",
		},
		{
			email:       "testemail8",
			password:    "testpassword8",
			statusCode:  404,
			description: "Should not find user with email",
		},
	}

	t.Log("TestUserFind")

	initAndResetDB()
	for i := 1; i <= 5; i++ {
		addUser(fmt.Sprintf("testemail%d", i), fmt.Sprintf("testpassword%d", i))
	}

	url := "/user/find"
	app := fiber.New()
	app.Get(url, controllers.UserFind)

	for _, test := range tests {
		values := map[string]string{"email": test.email}

		jsonValue, err := json.Marshal(values)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonValue))
		if err != nil {
			t.Fatal(err)
		}

		resp, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, test.statusCode, resp.StatusCode, test.description)
	}
}
