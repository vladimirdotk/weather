package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/vladimirdotk/weather/handlers"
)

// Run HTTP Service. Save coordinates, get weather forecast.
func Run(wg *sync.WaitGroup) {
	defer wg.Done()

	mux := http.NewServeMux()

	mux.HandleFunc("/coords", handlers.HandleCoords)
	mux.HandleFunc("/weather", handlers.HandleWeather)

	addr := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
	log.Fatal(http.ListenAndServe(addr, mux))
}
