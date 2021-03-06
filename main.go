package main

import (
	"log"
	"station-service/api"
	"station-service/config"
	"station-service/db"
)

func main() {

	// Load configuration settings.
	config, err := config.New(".")
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// Connect to the database.
	store, err := db.Connect(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Create a server and setup routes.
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("Failed to create a server: ", err)
	}

	// Start a server.
	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("Failed to start a server: ", err)
	}
}
