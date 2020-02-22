package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/bungysheep/news-api/pkg/configs"
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
	connString := os.Getenv(configs.DBCONNSTRINGVARIABLE)
	if connString != "" {
		return "", fmt.Errorf("Database connection string variable (%s) must be specified", configs.DBCONNSTRINGVARIABLE)
	}
	return connString, nil
}
