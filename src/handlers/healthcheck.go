package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetHealthCheck(w http.ResponseWriter, _ *http.Request) {
	healthMessage := map[string]string{"status": "OK"}

	jsonOut, _ := json.Marshal(healthMessage)
	if _, err := fmt.Fprintf(w, string(jsonOut)); err != nil {
		log.Fatalln("Error raised", err)
	}
}
