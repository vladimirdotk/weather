package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/vladimirdotk/weather/handlers"
)

// Run HTTP Service. Save coordinates, get weather forecast.
func Run(wg *sync.WaitGroup) {
	defer wg.Done()

	mux := http.NewServeMux()
	mux.HandleFunc("/coords", handlers.HandleCoords)
	mux.HandleFunc("/weather", handlers.HandleWeather)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")),
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe error: %v", err)
		}
	}()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	}

	log.Print("HTTP server received SIGINT, will shutdown gracefully")
}
