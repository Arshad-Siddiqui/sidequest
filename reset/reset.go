package reset

import (
	"github.com/arshad-siddiqui/sidequest/initialize"
	"github.com/arshad-siddiqui/sidequest/models"
)

// ResetDB resets the database
func ResetDB() {
	initialize.DB.Migrator().DropTable(&models.User{})
	initialize.DB.AutoMigrate(&models.User{})
}
