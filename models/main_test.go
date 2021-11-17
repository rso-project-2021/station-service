package models

import (
	"log"
	"os"
	"station-service/config"
	"station-service/db"
	"testing"
)

func TestMain(m *testing.M) {

	// Load configuration settings.
	config, err := config.LoadConfig("../.")
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	// Connect to the database.
	if err := db.Connect(config.DBDriver, config.DBSource); err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Run tests.
	code := m.Run()

	os.Exit(code)
}
