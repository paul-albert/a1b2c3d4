package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"goji.io"
	"goji.io/pat"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	"paul-albert-dd-technical-test/src/datasource"
	"paul-albert-dd-technical-test/src/handlers"
	"paul-albert-dd-technical-test/src/middleware"
)

const (
	ApiPrefix  = ""
	ListenAddr = "localhost:8080"
)

type App struct {
	EnvFileName string
	DB          *gorm.DB
	Log         bool
	Logger      *logrus.Logger
	Router      *goji.Mux
}

func (app *App) Initialize() {
	app.SetupLogger()

	app.InitializeEnv()

	DB, err := datasource.GetDB(app.Log, app.Logger)
	if err != nil {
		log.Fatalln(err)
	}

	app.DB = DB
	app.Router = goji.NewMux()

	app.AddMiddleware()
	app.InitializeRoutes()
}

func (app *App) SetupLogger() {
	appLogger := logrus.New()
	appLogger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:        "2006-01-02 15:04:05.999",
		FullTimestamp:          true,
		DisableTimestamp:       false,
		DisableLevelTruncation: true,
		ForceColors:            true,
		PadLevelText:           true,
	})
	appLogger.SetLevel(logrus.DebugLevel)
	appLogger.SetOutput(os.Stdout)

	app.Logger = appLogger
}

func (app *App) InitializeEnv() {
	if app.Log == true {
		app.Logger.Debug(
			fmt.Sprintf("\tApp: loading the '%s' file...", app.EnvFileName))
	}
	if err := godotenv.Load(app.EnvFileName); err != nil {
		app.Logger.Fatal(
			fmt.Sprintf("\tApp: error loading the '%s' file", app.EnvFileName))
	}

	if app.Log == true {
		app.Logger.Debug(
			fmt.Sprintf("\tApp: loaded the '%s' file.", app.EnvFileName))
	}
}

func (app *App) InitializeRoutes() {
	if app.Log == true {
		app.Logger.Debug("\tApp: initializing routes")
	}

	var apiPrefix = ApiPrefix
	if value, envFound := os.LookupEnv("API_PREFIX"); envFound != false {
		apiPrefix = value
	}

	routes := map[string]func(w http.ResponseWriter, _ *http.Request){
		// add API endpoints route for "health-check":
		"/healthcheck": handlers.GetHealthCheck,

		// add API endpoints routes for genres:
		"/genres":                handlers.GetAllGenres(app.DB),
		"/genres/generic":        handlers.GetAllGenericGenres(app.DB),
		"/genres/statistics":     handlers.GetGenresStatistics(app.DB),
		"/genres/search/:search": handlers.SearchGenresByName(app.DB),

		// add API endpoints routes for songs:
		"/songs":         handlers.GetAllSongs(app.DB),
		"/songs/generic": handlers.GetAllGenericSongs(app.DB),
		"/songs/search_by_length/:min_length/:max_length": handlers.SearchSongsByLengthsRange(app.DB),
		"/songs/search/:search":                           handlers.SearchSongsAndGenres(app.DB),
	}
	for route, handler := range routes {
		app.Router.HandleFunc(
			pat.Get(fmt.Sprintf("%s%s", apiPrefix, route)),
			handler)
	}
}

func (app *App) AddMiddleware() {
	if app.Log == true {
		app.Logger.Debug("\tApp: adding middleware")
	}

	app.Router.Use(middleware.JsonContentType)
	app.Router.Use(middleware.RequestLogging(app.Log, app.Logger))
}

func (app *App) Run() {
	if app.Log == true {
		app.Logger.Debug("\tApp: running")
	}

	var listenAddr = ListenAddr
	if value, envFound := os.LookupEnv("LISTEN"); envFound != false {
		listenAddr = value
	}

	if app.Log == true {
		app.Logger.Debug(fmt.Sprintf("\tApp: listening on %s\n", listenAddr))
		app.Logger.Debug("")
	}

	log.Fatalln(http.ListenAndServe(listenAddr, app.Router))
}
