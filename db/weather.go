package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/vladimirdotk/weather/types"
)

const (
	// InsertCoordsSQL Insert statement for saving latitude and longitude
	InsertCoordsSQL = "INSERT INTO WEATHER(LAT, LONG) VALUES($1,$2) RETURNING ID"
	// UpdateWeatherSQL Save weather data
	UpdateWeatherSQL = `
		UPDATE WEATHER
		SET TEMP=$1, HUMIDITY=$2, PRESSURE=$3
		WHERE ID = $4
	`
	// GetWeatherSQL Select weather data from storage
	GetWeatherSQL = `
		SELECT ID, LAT, LONG, TEMP, HUMIDITY, PRESSURE
		FROM WEATHER
		WHERE ID = $1
		AND TEMP IS NOT NULL
		AND HUMIDITY IS NOT NULL
		AND PRESSURE IS NOT NULL
		`
	// GetCoordsWithoutWeatherSQL Select ID and coords without weather
	GetCoordsWithoutWeatherSQL = `
		SELECT ID, LAT, LONG
		FROM WEATHER
		WHERE (
			TEMP IS NULL
			OR HUMIDITY IS NULL
			OR PRESSURE IS NULL
		)
	`
)

// GetWeather get weather from storage
func GetWeather(weather *types.Weather) (bool, error) {
	// Connect to db
	conn, err := connDb()
	if err != nil {
		return false, err
	}

	defer conn.Close(context.Background())

	// Get weather
	row := conn.QueryRow(context.Background(), GetWeatherSQL, weather.ID)
	err = weatherScan(&row, weather)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	log.Printf("Received weather from storage %+v", weather)

	return true, nil
}

// GetCoordsWithoutWeather get ID and coords without weather
func GetCoordsWithoutWeather() ([]*types.Weather, error) {
	// Connect to db
	conn, err := connDb()
	if err != nil {
		return nil, err
	}

	defer conn.Close(context.Background())

	// Get coords without weather
	rows, err := conn.Query(context.Background(), GetCoordsWithoutWeatherSQL)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var weatherRows []*types.Weather

	// Scan rows from db
	for rows.Next() {
		weatherRow := types.Weather{}
		err := rows.Scan(&weatherRow.ID, &weatherRow.Lat, &weatherRow.Long)
		if err != nil {
			return nil, err
		}
		weatherRows = append(weatherRows, &weatherRow)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return weatherRows, nil
}

// SaveCoords saves coordinates to storage, populates Weather structure with ID
func SaveCoords(weather *types.Weather) error {
	// Connect to db
	conn, err := connDb()
	if err != nil {
		return err
	}

	defer conn.Close(context.Background())

	// Insert coordinates
	err = conn.QueryRow(context.Background(), InsertCoordsSQL, weather.Lat, weather.Long).Scan(&weather.ID)
	if err != nil {
		return err
	}

	return nil
}

// SaveWeather saves weather data to storage
func SaveWeather(weatherData []*types.Weather) error {
	// Connect to db
	conn, err := connDb()
	if err != nil {
		return err
	}

	defer conn.Close(context.Background())

	for _, w := range weatherData {
		_, err := conn.Exec(context.Background(), UpdateWeatherSQL, w.Temp, w.Humidity, w.Pressure, w.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// Scan Row Data to Weather struct
func weatherScan(row *pgx.Row, w *types.Weather) error {
	return (*row).Scan(&w.ID, &w.Lat, &w.Long, &w.Temp, &w.Humidity, &w.Pressure)
}

// Connect to db
func connDb() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), getConnStr())

	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Get connection string
func getConnStr() string {
	db := os.Getenv("PG_DB")
	user := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PASS")
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, db)
}
