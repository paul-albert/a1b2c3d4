package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"paul-albert-dd-technical-test/src/app"
)

func AssertResponseCode(t *testing.T, expected, actual int) {
	assert.Equal(t, expected, actual,
		fmt.Sprintf("Expected response code %d, but got %d\n", expected, actual))
}

func PerformRequest(testApp app.App, request *http.Request) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	testApp.Router.ServeHTTP(responseRecorder, request)

	return responseRecorder
}
