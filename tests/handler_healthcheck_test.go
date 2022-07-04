package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandlerHealthcheck(t *testing.T) {
	request, _ := http.NewRequest(
		"GET",
		"/api/v1/healthcheck",
		nil)
	response := PerformRequest(testApp, request)

	AssertResponseCode(t, http.StatusOK, response.Code)

	expected, body := `{"status":"OK"}`, response.Body.String()
	assert.Equal(t, expected, body,
		fmt.Sprintf("Expected body %v, but got %v", expected, body))
}

func TestHandlerHealthcheckNotExistingUrl(t *testing.T) {
	request, _ := http.NewRequest(
		"GET",
		"/api/v1/healthcheck_not_existing_url",
		nil)
	response := PerformRequest(testApp, request)

	AssertResponseCode(t, http.StatusNotFound, response.Code)
}

func TestHandlerHealthcheckNotExistingBody(t *testing.T) {
	request, _ := http.NewRequest(
		"GET",
		"/api/v1/healthcheck",
		nil)
	response := PerformRequest(testApp, request)

	AssertResponseCode(t, http.StatusOK, response.Code)

	expected, body := `{"some_stuff":"BLA-BLA-BLA"}`, response.Body.String()
	assert.NotEqualf(t, expected, body,
		fmt.Sprintf("Expected body %v, but got %v", expected, body))
}
