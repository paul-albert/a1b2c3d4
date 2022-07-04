# Sample Golang project

## Goal

Build an API to deliver data (see description below) in JSON format.

* The API should make all songs searchable by artist, song or genre.
* The genre should be searchable by name in genres table.
* Returned data should include song, artist, genre name and length.
* Dependencies should be handled using Go modules.

## Data

Data are stored in SQLite database (the file `data.db`). It contains 2 tables with following structure:
* Songs
  * id (int)
  * artist (varchar(1024))
  * song (varchar(1024))
  * genre (int)
  * length (int)
* Genres
  * id (int)
  * name (varchar(32))

## Extra-functionality

* function in the API that returns songs by length, which need to search by passing of minimum and maximum length.
* function in the API that returns list of genres, number of songs and total length of all songs by genre. 
* unit testing

## Prerequisites:

* Installed Make tool
* Copy the file .env.sample to the new file .env in the root directory and check/edit values there

## How to

### Ensure that you have installed Go

> make go_version

### How to install all dependencies

> make dependencies

or

> make install

### How to compile

> make build

or

> make compile

### How to run

> make run

### How to test

> make test

or

> make tests

### How to clean up

> make clean

or

> make cleanup


## List of the API endpoints

* Health check (just to ensure that the server is live)

  `GET http://localhost:8080/api/v1/healthcheck`

* List of genres (together with their songs)

  `GET http://localhost:8080/api/v1/genres`

* List of genres (generic list, without songs)

  `GET http://localhost:8080/api/v1/genres/generic`

* Information about genres statistics (list of genres and number of songs and total length of all songs by genre)

  `GET http://localhost:8080/api/v1/genres/statistics`

* Search and output genres by name

  `GET http://localhost:8080/api/v1/genres/search/:search`
  
  * Sample: `GET http://localhost:8080/api/v1/genres/search/rock`

* List of songs (together with their genres)

  `GET http://localhost:8080/api/v1/songs`

* List of songs (generic list, without genres)

  `GET http://localhost:8080/api/v1/songs/generic`

* Search songs by length, i.e. search by passing of minimum and maximum lengths

  `GET http://localhost:8080/api/v1/songs/search_by_length/:min_length/:max_length`

  * Sample: `GET http://localhost:8080/api/v1/songs/search_by_length/100/200`

* Search songs by artist, song or genre

  `GET http://localhost:8080/api/v1/songs/search/:search`
  
  * Sample: `GET http://localhost:8080/api/v1/songs/search/beatles`

## Used software/tools

* Github/git
* make - MinGW on Windows or GNU Make (built-in) on MacOS
* Windows/MacOS
* Google and Stackoverflow are best friends of developer :)

## Further suggestions

#### About the structure of the data

* Fields `artist` and `song` in table `songs` are better to make as `VARCHAR(256)`, 256 is quite enough - 1024 is too big
* Rename the field `songs`.`genre` to `songs`.`genre_id` - it's more clear and readable
* Maybe, the field `songs`.`genre` is better make as ENUM type (not sure about the SQLite engine) with values from the `genres` table, then the `genres` table is not necessary

#### What fields would you index in these tables

* Field `genres`.`name`:
  * add UNIQUE INDEX
* Field `songs`.`artist` and `songs`.`song`:
  * add composite UNIQUE INDEX for the both fields (as combination of artist and song should be unique by idea)
* Field `songs`.`genre`:
  * add INDEX
  * add FOREIGN KEY constraint (maybe, with possibility of CASCADE update/deletion)
* Fields `genres`.`name`, `songs`.`artist`, `songs`.`song` - would be good to index as FULLTEXT SEARCH (FTS)  
  * SQLite has extension FTS5 extension
  * Other engines (like as Postgres or MySQL) has built-in support for the FTS
  * Also, there are external solutions: Elasticsearch, Apache Lucene/Solr, Xapian, Sphinx etc.
