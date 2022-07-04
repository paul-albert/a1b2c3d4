package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Song struct {
	ID      int    `gorm:"primary_key, AUTO_INCREMENT, column:ID" json:"id"`
	Artist  string `gorm:"column:artist" json:"artist"`
	Song    string `gorm:"column:song" json:"song"`
	GenreID int    `gorm:"column:genre" json:"genre_id"`
	Length  int    `gorm:"column:length" json:"length"`
	// relationship 1:1
	Genre GenericGenre
}

type GenericSong struct {
	ID      int    `gorm:"primary_key, AUTO_INCREMENT, column:ID" json:"id"`
	Artist  string `gorm:"column:artist" json:"artist"`
	Song    string `gorm:"column:song" json:"song"`
	GenreID int    `gorm:"column:genre" json:"genre_id"`
	Length  int    `gorm:"column:length" json:"length"`
	// (here the json:"-" means to be omitted in JSON-"output"!)
	Genre Genre `gorm:"ForeignKey:genre" json:"-"`
}

func (song *Song) TableName() string {
	return "songs"
}

func (song *GenericSong) TableName() string {
	return "songs"
}

func (song Song) ToString() string {
	return fmt.Sprintf(
		"Song: { id=%d, artist=\"%s\", song=\"%s\", genre_id=%d, length=%d }",
		song.ID, song.Artist, song.Song, song.GenreID, song.Length)
}

type SongModel struct {
}

type SongGenericModel struct {
}

type SongsAndGenresSearchResult struct {
	Song   string
	Artist string
	Genre  string
	Length int
}

func (songModel SongModel) FindAll(db *gorm.DB) ([]Song, error) {
	var songs []Song
	db.Preload("Genre").Find(&songs)
	return songs, nil
}

func (songModel SongGenericModel) FindAll(db *gorm.DB) ([]GenericSong, error) {
	var songs []GenericSong
	db.Preload("Genre").Find(&songs)
	return songs, nil
}

func (songModel SongGenericModel) FindByLengthsRange(
	db *gorm.DB, minLength int, maxLength int) ([]GenericSong, error) {

	var songs []GenericSong
	db.Where("length >= ? AND length <= ?", minLength, maxLength).Find(&songs)
	return songs, nil
}

func (songModel SongModel) Search(db *gorm.DB, search string) ([]SongsAndGenresSearchResult, error) {
	/*
		// The SQL query should be something like:
		SELECT
			songs.artist AS artist,
			songs.song AS song,
			genres.name AS genre,
			songs.length AS length
		FROM songs
			INNER JOIN genres ON songs.genre = genres.ID
		WHERE
			songs.artist LIKE "%some_search_string%"
			OR songs.song LIKE "%some_search_string%"
			OR genres.name LIKE "%some_search_string%";
	*/
	var results []SongsAndGenresSearchResult
	// The double "%%" in the "%%%s%%" means escape of literal percent sign
	var searchStr = fmt.Sprintf("%%%s%%", search)
	db.
		Model(&Song{}).
		Where("songs.artist LIKE ?", searchStr).
		Or("songs.song LIKE ?", searchStr).
		Or("genres.name LIKE ?", searchStr).
		Select(
			"songs.artist AS artist, " +
				"songs.song AS song, " +
				"genres.name AS genre, " +
				"songs.length AS length").
		Joins("INNER JOIN genres ON songs.genre = genres.ID").
		Find(&results)
	return results, nil
}
