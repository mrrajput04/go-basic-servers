package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver import with blank identifier

	"github.com/joho/godotenv"
)

// DB is a global variable to hold the database connection.
var DB *sql.DB

// Connect initializes the database connection using environment variables.
func ConnectDB() {
	// Load environment variables from the .env file.
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env files")
	}

	// Retrieve database connection details from environment variables.
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Create a connection string using retrieved details.

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open a new database connection.
	db, err := sql.Open("mysql", connectionString)

	// Check for errors during database connection.
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// Set the global DB variable to the connected database.
	DB = db
}
