package handlers

import (
	"encoding/json"
	"fmt"
	"goji.io/pat"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"

	"paul-albert-dd-technical-test/src/models"
)

func GetAllSongs(db *gorm.DB) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		var songModel models.SongModel
		songs, _ := songModel.FindAll(db)
		jsonOut, _ := json.Marshal(songs)
		if _, err := fmt.Fprintf(w, string(jsonOut)); err != nil {
			log.Fatalln("Error raised", err)
		}
	}
}

func GetAllGenericSongs(db *gorm.DB) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		var songModel models.SongGenericModel
		songs, _ := songModel.FindAll(db)
		jsonOut, _ := json.Marshal(songs)
		if _, err := fmt.Fprintf(w, string(jsonOut)); err != nil {
			log.Fatalln("Error raised", err)
		}
	}
}

func SearchSongsByLengthsRange(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: make some validation for the input parameters:
		// TODO: 	cast both min_length and max_length as string;
		// TODO: 	min_length < max_length;
		// TODO: 	0 < min_length < INF;
		// TODO: 	0 < max_length < INF;

		strMinLength := pat.Param(r, "min_length")
		strMaxLength := pat.Param(r, "max_length")
		minLength, _ := strconv.Atoi(strMinLength)
		maxLength, _ := strconv.Atoi(strMaxLength)

		var songModel models.SongGenericModel
		songs, _ := songModel.FindByLengthsRange(db, minLength, maxLength)
		jsonOut, _ := json.Marshal(songs)
		if _, err := fmt.Fprintf(w, string(jsonOut)); err != nil {
			log.Fatalln("Error raised", err)
		}
	}
}

func SearchSongsAndGenres(db *gorm.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: maybe, make some validation for the input parameter:
		// TODO: 	length of search string must be > 0;

		search := pat.Param(r, "search")

		var songModel models.SongModel
		genres, _ := songModel.Search(db, search)
		jsonOut, _ := json.Marshal(genres)
		if _, err := fmt.Fprintf(w, string(jsonOut)); err != nil {
			log.Fatalln("Error raised", err)
		}
	}
}
