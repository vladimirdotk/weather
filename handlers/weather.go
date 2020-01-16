package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/vladimirdotk/weather/db"
	"github.com/vladimirdotk/weather/response"
	"github.com/vladimirdotk/weather/types"
)

// HandleWeather gets ID, gets weather data from storage, sends to client
func HandleWeather(w http.ResponseWriter, r *http.Request) {
	// Validate HTTP Method
	if r.Method != "GET" {
		err := map[string]string{"errors": "Only GET method is allowed"}
		response.SendJSON(w, err, http.StatusMethodNotAllowed)
		return
	}

	// Get ID
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.SendJSON(w, map[string]string{"errors": "You should provide a valid id"}, http.StatusBadRequest)
		return
	}

	weather := &types.Weather{}
	weather.ID = id

	// Get Weather
	found, err := db.GetWeather(weather)
	if err != nil {
		log.Printf("Error getting weather %s", err.Error())
		response.SendJSON(w, map[string]string{"errors": "Internal Server Error"}, http.StatusInternalServerError)
		return
	}

	// Send 404 if not found
	if !found {
		response.SendJSON(w, map[string]string{"errors": "Weather not found"}, http.StatusNotFound)
		return
	}

	response.SendJSON(w, weather, http.StatusOK)
}
