package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGenres(t *testing.T) {
	request, _ := http.NewRequest(
		"GET",
		"/api/v1/genres",
		nil)
	response := PerformRequest(testApp, request)

	AssertResponseCode(t, http.StatusOK, response.Code)

	expected := `[{"id":1,"name":"rock","Songs":[{"id":1,"artist":"Artist_1","song":"Song_1","genre_id":1,"length":123},{"id":2,"artist":"Artist_2","song":"Song_2","genre_id":1,"length":123},{"id":3,"artist":"Artist_3","song":"Song_3","genre_id":1,"length":123}]},{"id":2,"name":"jazz","Songs":[{"id":4,"artist":"Artist_1","song":"Song_1","genre_id":2,"length":123},{"id":5,"artist":"Artist_2","song":"Song_2","genre_id":2,"length":123},{"id":6,"artist":"Artist_3","song":"Song_3","genre_id":2,"length":123}]}]`
	body := response.Body.String()
	assert.Equal(t, expected, body,
		fmt.Sprintf("Expected body %v, but got %v", expected, body))
}

func TestHandlerGenresGeneric(t *testing.T) {
	request, _ := http.NewRequest(
		"GET",
		"/api/v1/genres/generic",
		nil)
	response := PerformRequest(testApp, request)

	AssertResponseCode(t, http.StatusOK, response.Code)

	expected := `[{"id":1,"name":"rock"},{"id":2,"name":"jazz"}]`
	body := response.Body.String()
	assert.Equal(t, expected, body,
		fmt.Sprintf("Expected body %v, but got %v", expected, body))
}

func TestGenresStatistics(t *testing.T) {
	request, _ := http.NewRequest(
		"GET",
		"/api/v1/genres/statistics",
		nil)
	response := PerformRequest(testApp, request)

	AssertResponseCode(t, http.StatusOK, response.Code)

	expected := `[{"ID":2,"Name":"jazz","SongsCount":3,"TotalLength":369},{"ID":1,"Name":"rock","SongsCount":3,"TotalLength":369}]`
	body := response.Body.String()
	assert.Equal(t, expected, body,
		fmt.Sprintf("Expected body %v, but got %v", expected, body))
}

func TestGenresSearch1(t *testing.T) {
	request, _ := http.NewRequest(
		"GET",
		"/api/v1/genres/search/rock",
		nil)
	response := PerformRequest(testApp, request)

	AssertResponseCode(t, http.StatusOK, response.Code)

	expected := `[{"id":1,"name":"rock"}]`
	body := response.Body.String()
	assert.Equal(t, expected, body,
		fmt.Sprintf("Expected body %v, but got %v", expected, body))
}

func TestGenresSearch2(t *testing.T) {
	request, _ := http.NewRequest(
		"GET",
		"/api/v1/genres/search/jazz",
		nil)
	response := PerformRequest(testApp, request)

	AssertResponseCode(t, http.StatusOK, response.Code)

	expected := `[{"id":2,"name":"jazz"}]`
	body := response.Body.String()
	assert.Equal(t, expected, body,
		fmt.Sprintf("Expected body %v, but got %v", expected, body))
}

func TestGenresSearch3(t *testing.T) {
	request, _ := http.NewRequest(
		"GET",
		"/api/v1/genres/search/rap",
		nil)
	response := PerformRequest(testApp, request)

	AssertResponseCode(t, http.StatusOK, response.Code)

	expected := `[]`
	body := response.Body.String()
	assert.Equal(t, expected, body,
		fmt.Sprintf("Expected body %v, but got %v", expected, body))
}
