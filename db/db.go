package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func Connect(source, driver string) (err error) {

	// Connect to database.
	db, err = sqlx.Connect(source, driver)
	if err != nil {
		return
	}

	// Test database connection.
	if err = db.Ping(); err != nil {
		return
	}

	log.Println("Connected to database!")
	return
}

func GetDB() *sqlx.DB {
	return db
}
