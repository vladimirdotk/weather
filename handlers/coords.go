package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vladimirdotk/weather/db"
	"github.com/vladimirdotk/weather/response"
	"github.com/vladimirdotk/weather/types"
	"github.com/vladimirdotk/weather/validators"
)

// HandleCoords get coordinates, write to storage, return ID
func HandleCoords(w http.ResponseWriter, r *http.Request) {
	// Validate HTTP Method
	if r.Method != "POST" {
		err := map[string]string{"errors": "Only POST method is allowed"}
		response.SendJSON(w, err, http.StatusMethodNotAllowed)
		return
	}

	weather := &types.Weather{}

	// Decode data
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(weather); err != nil {
		log.Printf("Error decoding coords %s", err.Error())
		response.SendJSON(w, map[string]string{"errors": "Bad request"}, http.StatusBadRequest)
		return
	}

	// Validate data
	if validErrors := validators.ValidateCoords(weather); len(validErrors) > 0 {
		err := map[string]map[string][]string{"errors": validErrors}
		response.SendJSON(w, err, http.StatusBadRequest)
		return
	}

	// Save coordinates to storage
	if err := db.SaveCoords(weather); err != nil {
		log.Printf("Error saving coords %s", err.Error())
		response.SendJSON(w, map[string]string{"errors": "Internal Server Error"}, http.StatusInternalServerError)
		return
	}

	response.SendJSON(w, weather, http.StatusOK)
}
