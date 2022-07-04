package tests

import (
	"fmt"

	"paul-albert-dd-technical-test/src/app"
	"paul-albert-dd-technical-test/src/models"
)

func Setup(testApp app.App) {
	if err := testApp.DB.AutoMigrate(&models.Genre{}, &models.Song{}); err != nil {
		testApp.Logger.Fatalln("Couldn't create DB tables.")
		return
	} else {
		testApp.Logger.Debug("DB tables created.")
	}

	genres := []string{"rock", "jazz"}
	for _, i := range []int{0, 1} {
		genre := models.Genre{Name: genres[i]}
		genreResult := testApp.DB.Create(&genre)
		testApp.Logger.Debug(fmt.Sprintf("\tInserted genres: %d rows", genreResult.RowsAffected))
		testApp.Logger.Debug(fmt.Sprintf("\tInserted genre id: %d", genre.ID))

		for _, j := range []int{1, 2, 3} {
			song := models.Song{
				Artist:  fmt.Sprintf("Artist_%d", j),
				Song:    fmt.Sprintf("Song_%d", j),
				GenreID: genre.ID,
				Length:  123,
			}
			songResult := testApp.DB.Create(&song)
			testApp.Logger.Debug(fmt.Sprintf("\tInserted songs: %d rows", songResult.RowsAffected))
			testApp.Logger.Debug(fmt.Sprintf("\tInserted song id: %d", song.ID))
		}
	}

	fmt.Println("Unit tests: Setup completed.")
}
