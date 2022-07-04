package handlers

import (
	"encoding/json"
	"fmt"
	"goji.io/pat"
	"gorm.io/gorm"
	"log"
	"net/http"

	"paul-albert-dd-technical-test/src/models"
)

func GetAllGenres(db *gorm.DB) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		var genreModel models.GenreModel
		genres, _ := genreModel.FindAll(db)
		jsonOut, _ := json.Marshal(genres)
		if _, err := fmt.Fprintf(w, string(jsonOut)); err != nil {
			log.Fatalln("Error raised", err)
		}
	}
}

func GetAllGenericGenres(db *gorm.DB) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		var genreModel models.GenreGenericModel
		genres, _ := genreModel.FindAll(db)
		jsonOut, _ := json.Marshal(genres)
		if _, err := fmt.Fprintf(w, string(jsonOut)); err != nil {
			log.Fatalln("Error raised", err)
		}
	}
}

func GetGenresStatistics(db *gorm.DB) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		var genreModel models.GenreModel
		statistics, _ := genreModel.GetGenresStatistics(db)
		jsonOut, _ := json.Marshal(statistics)
		if _, err := fmt.Fprintf(w, string(jsonOut)); err != nil {
			log.Fatalln("Error raised", err)
		}
	}
}

func SearchGenresByName(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: maybe, make some validation for the input parameter:
		// TODO: 	length of search string must be > 0;

		search := pat.Param(r, "search")

		var genreModel models.GenreGenericModel
		genres, _ := genreModel.Search(db, search)
		jsonOut, _ := json.Marshal(genres)
		if _, err := fmt.Fprintf(w, string(jsonOut)); err != nil {
			log.Fatalln("Error raised", err)
		}
	}
}
