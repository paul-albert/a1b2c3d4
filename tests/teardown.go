package tests

import (
	"fmt"

	"paul-albert-dd-technical-test/src/app"
	"paul-albert-dd-technical-test/src/models"
)

func Teardown(testApp app.App) {
	if err := testApp.DB.Migrator().DropTable(&models.Song{}, &models.Genre{}); err != nil {
		testApp.Logger.Fatalln("Couldn't drop DB tables.")
	} else {
		testApp.Logger.Debug("DB tables dropped.")
	}

	fmt.Println("Unit tests: Teardown completed.")
}
