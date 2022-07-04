package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Genre struct {
	ID   int    `gorm:"primary_key, AUTO_INCREMENT, column:ID" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	// relationship 1:N
	Songs []GenericSong `gorm:"ForeignKey:genre"`
}

type GenericGenre struct {
	ID   int    `gorm:"primary_key, AUTO_INCREMENT, column:ID" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	// (here the json:"-" means to be omitted in JSON-"output"!)
	Songs []Song `gorm:"ForeignKey:genre" json:"-"`
}

func (genre *Genre) TableName() string {
	return "genres"
}

func (genre *GenericGenre) TableName() string {
	return "genres"
}

func (genre Genre) ToString() string {
	return fmt.Sprintf(
		"Genre: { id=%d, name=\"%s\" }",
		genre.ID, genre.Name)
}

type GenreModel struct {
}

type GenreGenericModel struct {
}

type GenresStatisticsResult struct {
	ID          int
	Name        string
	SongsCount  int
	TotalLength int
}

func (genreModel GenreModel) FindAll(db *gorm.DB) ([]Genre, error) {
	var genres []Genre
	db.Preload("Songs").Find(&genres)
	return genres, nil
}

func (genreModel GenreGenericModel) FindAll(db *gorm.DB) ([]GenericGenre, error) {
	var genres []GenericGenre
	db.Find(&genres)
	return genres, nil
}

func (genreModel GenreModel) GetGenresStatistics(db *gorm.DB) ([]GenresStatisticsResult, error) {
	/*
		// The SQL query should be something like:
		SELECT
			genres.ID, genres.name,
			COUNT(songs.iD) as songs_count,
			IFNULL(SUM(songs.length), 0) as total_length
		FROM genres
			LEFT JOIN songs ON genres.ID = songs.genre
		GROUP BY genres.ID
		ORDER BY genres.name;
	*/
	var results []GenresStatisticsResult
	db.
		Model(&Genre{}).
		Select(
			"genres.ID, " +
				"genres.name, " +
				"COUNT(songs.iD) as songs_count, " +
				"IFNULL(SUM(songs.length), 0) as total_length").
		Joins("LEFT JOIN songs ON genres.ID = songs.genre").
		Group("genres.ID").
		Order("genres.name").
		Find(&results)
	return results, nil
}

func (genreModel GenreGenericModel) Search(db *gorm.DB, search string) ([]GenericGenre, error) {
	/*
		// The SQL query should be something like:
		SELECT
			id,
			name
		FROM genres
		WHERE name LIKE "%some_search_string%";
	*/
	var genres []GenericGenre
	// The double "%%" in the "%%%s%%" means escape of literal percent sign
	var searchStr = fmt.Sprintf("%%%s%%", search)
	db.
		Model(&Genre{}).
		Where("name LIKE ?", searchStr).
		Select("ID", "name").
		Find(&genres)
	return genres, nil
}
