package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

func InitDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Default values
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "1433"
	}
	if user == "" {
		user = "sa"
	}
	if dbname == "" {
		dbname = "BlogDB"
	}

	// Connection string for SQL Server
	connString := fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=%s;encrypt=disable;TrustServerCertificate=true",
		host, port, user, password, dbname)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	return db, nil
}