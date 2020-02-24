package database

import (
	"database/sql"
	"log"
	"net"
	"os"
	"time"

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

	for i := 0; i < configs.NUMBERDIALATTEMPT; i++ {
		err = DbConnection.Ping()
		if err != nil {
			opErr, ok := err.(*net.OpError)
			if !ok || opErr.Op != "dial" {
				return err
			}
			time.Sleep(5 * time.Second)
		}
	}

	return err
}

func resolveDbConnectionString() (string, error) {
	connString := os.Getenv("DATABASE_URL")
	if connString != "" {
		return connString, nil
	}

	return configs.DBCONNSTRING, nil
}
