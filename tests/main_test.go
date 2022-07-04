package tests

import (
	"fmt"
	"os"
	"testing"

	"paul-albert-dd-technical-test/src/app"
)

var testApp app.App

func TestMain(m *testing.M) {
	testApp = app.App{
		EnvFileName: "./../.env.test",
		Log:         false,
	}
	testApp.Initialize()
	// NB: Here we don't need to call the `testApp.Run()`

	Setup(testApp)

	fmt.Println("Unit tests: Begin the running...")
	code := m.Run()
	fmt.Println("Unit tests: Finished the running.")

	Teardown(testApp)

	os.Exit(code)
}
