# List of TODO:

### (The points are open for editing)

* ~~Move DB file to "data" directory~~
* ~~Make the .gitignore file more "typical"~~
* ~~$GOPATH env variable - figure out for Windows/MacOS~~
* ~~Find out how to install/setup all necessary Go packages (like as Python's requirements.txt file) in some simple way~~
    * ~~"Virtual environment" like for Python (not necessary)~~
* ~~Divide the project into packages/subpackages (?), structure of project's folder(s) and file(s)~~
    * ~~"vendor" (not necessary)~~
    * ~~"cmd" (not necessary)~~
    * ~~"bin"~~
    * ~~"src"~~
    * ~~"data"~~
    * ~~"tests"~~
* ~~Makefile~~
  * ~~Find out a clean way to uninstall and install all dependencies from scratch~~
  * ~~Check all commands~~
* ~~Logging (some logger with different levels: info, debug etc.)~~
  * ~~Use the "https://github.com/apsdehal/go-logger" instead of logrus~~  
  * ~~Log all requests (method, URL, date, IP etc.)~~
* ~~DB:~~
  * ~~Make connection to DB~~
  * ~~Configure connection pool (?)~~
    * ~~By default, the GORM's connection pool works fine~~
  * ~~Move call of `datasource.GetDB()` from `models/*.go` to some other place/layer~~  
  * ~~Make setup of ORM~~
  * ~~Write the models `songs` and `genres`~~
    * ~~Relationships (one-to-many, 1:N, and one-to-one, 1:1)~~
    * ~~Definitions for GORM and JSON~~
  * ~~Log all SQL queries~~
    * ~~Useful option for single query: Debug() in chaining of db methods.~~
    * ~~Maybe, use some custom logger (as described in GORM's docs)~~
* ~~Figure out with the Jetbrains warnings for main.go~~
* ~~Move some hardcoded stuff (some configurations) to separate files:~~
  * ~~.env~~ 
  * ~~"yaml" (not necessary)~~
* ~~Implement functionality:~~
  * ~~*(Mandatory points)*:~~
    * ~~Search and output songs by artist, song or genre~~
      * ~~Full-text search (?) (no, enough to use the SELECT LIKE construction)~~
      * ~~JSON-data in output should include: song, artist, genre name and length~~
    * ~~Search and output genres by name~~
      * ~~Full-text search (?) (no, enough to use the SELECT LIKE construction)~~
      * ~~JSON-data in output should include: ID, name~~
  * ~~*(Extra-credit points)*:~~
    * ~~Return list of all genres~~
      * ~~Generic genres (without songs)~~
      * ~~Genres and their songs~~
    * ~~Return list of all songs~~
      * ~~Generic songs (without genres)~~
      * ~~Songs and their genres~~
    * ~~Return songs by length, which need to search by passing of minimum and maximum lengths~~
    * ~~Return list of genres and number of songs and total length of all songs by genre~~
* ~~Check HTTP-headers for responses:~~
  * ~~"Content-type" value (it would be "application/json" instead of "text/plain" by default)~~
* ~~*(Extra-credit)* Unit tests:~~
  * ~~Separate directory "tests"~~
  * ~~Generate/recreate/cleanup test data, in some separate testing DB (via the GORM)~~
    * ~~Create the test DB (SQLite in memory instead of DB on disk)~~
      * ~~GORM recommends to use "file::memory:?cache=shared"~~
      * ~~StackOverflow's recommendation: "file:some_db_name?mode=memory&cache=shared"~~
    * ~~Drop DB tables if exist (not necessary)~~
    * ~~Create DB tables on base of models (DB migrations)~~
    * ~~The 2 above points could be made on idea of DB migrations (GORM has own migrator)~~
    * ~~Fill DB tables with some test data (see GORM documentation)~~
    * ~~Drop the test DB (not necessary)~~
  * ~~Apparently provide way of setup/teardown, test suite (?)~~
  * ~~Own testing case for every handler + create the app instance (?)~~
* Validation of request parameters (?) (see TODO inside the code)
* ~~Date/time format for logger (+ milliseconds/microseconds)~~
* ~~Prepare Postman collection of API-endpoints~~
* ~~Think about:~~
  * ~~Suggestions regarding the structure of the data~~
  * ~~Which fields (and how) would be indexed in DB~~
