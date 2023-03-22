package controllers_test

import (
	. "github.com/arshad-siddiqui/sidequest/initialize"
	"github.com/arshad-siddiqui/sidequest/models"
	"github.com/arshad-siddiqui/sidequest/reset"
)

func addUser(email, password string) {
	user := models.User{Email: email, Password: password}
	DB.Create(&user)

}

func initAndResetDB() {
	initDB()
	reset.ResetDB()
}

func initDB() {
	LoadEnv("../.env.test")
	ConnectDB()
}
