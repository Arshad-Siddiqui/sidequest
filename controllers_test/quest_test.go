package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arshad-siddiqui/sidequest/controllers"
	"github.com/arshad-siddiqui/sidequest/initialize"
	"github.com/arshad-siddiqui/sidequest/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestQuestCreate(t *testing.T) {
	t.Log("TestQuestCreate")

	initAndResetDB()

	url := "/quest/create"
	app := fiber.New()
	app.Post(url, controllers.QuestCreate)

	// creating a quest
	quest := models.Quest{
		Title:       "Test Quest",
		Description: "Test Description",
	}
	body, _ := json.Marshal(quest)
	req := httptest.NewRequest(http.MethodPost, "/quest/create", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// checking if the quest was created
	var questFromDB models.Quest
	initialize.DB.Where("title = ?", quest.Title).First(&questFromDB)
	assert.Equal(t, quest.Title, questFromDB.Title)
	assert.Equal(t, quest.Description, questFromDB.Description)
}
