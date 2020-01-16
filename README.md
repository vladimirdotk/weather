# Weather

> Microservice to get weather by latitude and longitude

## Usage

+ `cp .env.example .env` and change variables
+ run `docker-compose up` or `go run main.go` (you will need to execute `./docker/db/init/init.sql` manually in case of not using db via docker-compose)
+ send `POST http://localhost:${HTTP_PORT}/coords` with `lat` and `long` JSON fields, get ID back
+ send `GET http://localhost:${HTTP_PORT}/weather?id={$id}` with ID from previous step and get weather back