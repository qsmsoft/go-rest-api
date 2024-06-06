package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %q", err)
	}

	createEventsTables()

}

func createEventsTables() {
	query := `
    CREATE TABLE IF NOT EXISTS events (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
		dateTime TIMESTAMP NOT NULL,
		user_id INTEGER
    	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Error creating tables: %q", err)
	}
}