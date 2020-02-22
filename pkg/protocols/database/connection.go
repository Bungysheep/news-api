package database

import (
	"database/sql"
	"github.com/bungysheep/news-api/pkg/configs"
	"log"
	"os"
)

var (
	// DbConnection - Database connection
	DbConnection *sql.DB
)

// CreateDbConnection - Creates connection to database
func CreateDbConnection() error {
	log.Printf("Creating database connection...")

	dbConnString, err := resolveDbConnectionString()
	if err != nil {
		return err
	}

	db, err := sql.Open("postgres", dbConnString)
	if err != nil {
		return err
	}

	DbConnection = db

	return DbConnection.Ping()
}

func resolveDbConnectionString() (string, error) {
	connString := os.Getenv("DATABASE_URL")
	if connString != "" {
		return connString, nil
	}

	return configs.DBCONNSTRING, nil
}
