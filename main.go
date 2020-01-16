package main

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/vladimirdotk/weather/api"
	"github.com/vladimirdotk/weather/worker"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env found")
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go api.Run(&wg)
	go worker.Run(&wg)

	wg.Wait()
}
