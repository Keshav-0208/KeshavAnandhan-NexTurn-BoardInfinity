package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func InitDB() error {
	var err error
	Database, err = sql.Open("sqlite3", "./blogs.db")
	if err != nil {
		return fmt.Errorf("unable to connect to the database: %v", err)
	}

	if err := Database.Ping(); err != nil {
		return fmt.Errorf("failed to establish a database connection: %v", err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS blog_posts (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		author TEXT NOT NULL,
		created_at TEXT NOT NULL
	);
	`
	_, err = Database.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("error creating blog posts table: %v", err)
	}

	log.Println("Database connected and schema is set up successfully.")
	return nil
}

func GetDatabase() *sql.DB {
	if Database == nil {
		log.Fatal("Database is not initialized.")
	}
	return Database
}
