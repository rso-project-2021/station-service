package main

import (
	"log"
	"station-service/config"
	"station-service/db"
	"station-service/server"
)

func main() {

	// Load configuration settings.
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// Connect to the database.
	if err := db.Connect(config.DBDriver, config.DBSource); err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	server.Start(config.ServerAddress, config.GinMode)
}
