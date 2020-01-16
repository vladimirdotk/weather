package worker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/vladimirdotk/weather/db"
	"github.com/vladimirdotk/weather/types"
)

const (
	urlPattern = "https://api.weatherbit.io/v2.0/current?lat=%s&lon=%s&lon&key=%s"
)

// Run worker logic periodically
func Run(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		process()
		time.Sleep(time.Minute)
	}
}

// Get weather from api and save to storage
func process() {
	log.Print("Running worker")

	weatherData, err := db.GetCoordsWithoutWeather()
	if err != nil {
		log.Printf("Error getting coords without weather %s", err.Error())
		return
	}

	if len(weatherData) == 0 {
		log.Print("No data for getting weather")
		return
	}

	getWeather(weatherData)
	db.SaveWeather(weatherData)
}

// Get weather from api to weather struct
func getWeather(weatherData []*types.Weather) {
	for _, weather := range weatherData {
		url := fmt.Sprintf(urlPattern, weather.Lat, weather.Long, os.Getenv("WEATHER_API_KEY"))

		// Make request for weather
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("Error api request by lat %s and long %s", weather.Lat, weather.Long)
		}

		defer resp.Body.Close()

		// Decode api response
		var apiResponse types.ApiResponse
		err = json.NewDecoder(resp.Body).Decode(&apiResponse)
		if err != nil {
			fmt.Print(url)
			log.Printf("Error decoding api response %s", err.Error())
		}

		log.Printf("Got data from api %+v", apiResponse)

		mapResponse(&apiResponse, weather)
	}
}

// Map api response to weather struct
func mapResponse(resp *types.ApiResponse, weather *types.Weather) {
	if resp.Count > 0 {
		weather.Temp = resp.Data[0].Temp
		weather.Humidity = resp.Data[0].Humidity
		weather.Pressure = resp.Data[0].Pressure
	}
}
