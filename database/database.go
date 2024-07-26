package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase() {
    var err error
    DB, err = sql.Open("sqlite3", "test.db")
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}
    log.Println("Database connection successfully opened")
    createTables()
}
func createTables() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT,
			password TEXT NOT NULL,
			created_at DATETIME,
			updated_at DATETIME
		);`,
		`CREATE TABLE IF NOT EXISTS recipes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			from_by INTEGER NOT NULL,
			description TEXT NOT NULL,
			value INTEGER NOT NULL,
			received_at DATETIME,
			received_by INTEGER NOT NULL,
			created_at DATETIME,
			updated_at DATETIME
		);`,
		`CREATE TABLE IF NOT EXISTS expenses (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			from_by INTEGER NOT NULL,
			description TEXT NOT NULL,
			value INTEGER NOT NULL,
			paid_at DATETIME,
			received_by INTEGER NOT NULL,
			created_at DATETIME,
			updated_at DATETIME
		);`,
	}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatal("Failed to create table:", err)
		}
	}

	log.Println("Database tables created")
}
func CloseDatabase() {
	err := DB.Close()
	if err != nil {
		log.Println("Failed to close database connection:", err)
	} else {
		log.Println("Database connection successfully closed")
	}
}