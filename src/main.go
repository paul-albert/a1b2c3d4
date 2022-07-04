package main

import (
	"paul-albert-dd-technical-test/src/app"
)

func main() {
	mainApp := app.App{
		EnvFileName: ".env",
		Log:         true,
	}
	mainApp.Initialize()
	mainApp.Run()
}
