package datasource

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetDB(appLog bool, appLogger *logrus.Logger) (*gorm.DB, error) {
	dbLogger := logger.New(
		// io writer
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		// logging config
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Enable color
		},
	)

	var dbPath = ""
	if value, envFound := os.LookupEnv("DBPATH"); envFound != false {
		dbPath = value
	}
	if dbPath != "" {
		dbLog, err := strconv.ParseBool(os.Getenv("DBLOG"))
		if err != nil {
			return nil, err
		}

		dbConfig := gorm.Config{}
		if dbLog != false {
			dbConfig = gorm.Config{
				Logger: dbLogger,
			}
		}

		if appLog == true {
			appLogger.Debug(
				fmt.Sprintf("\tDB: connecting to '%s'...", dbPath))
		}
		db, err := gorm.Open(sqlite.Open(dbPath), &dbConfig)
		if err != nil {
			return nil, err
		}

		if appLog == true {
			appLogger.Debug(
				fmt.Sprintf("\tDB: connected to '%s'", dbPath))
		}

		return db, err
	} else {
		return nil, errors.New("empty path to database")
	}
}
