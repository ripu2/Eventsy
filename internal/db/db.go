package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api/api.db")
	if err != nil {
		panic("Failed to connect to database!!!!")
	}
	fmt.Println("Database initialized successfully!!!")
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT NOT NULL,
    description TEXT NOT NULL,
		location TEXT NOT NULL,
    dateTime DATETIME NOT NULL,
		userId INTEGER NOT NULL
	)
	`
	_, err := DB.Exec(createEventTable)
	if err != nil {
		panic("Failed to create EventsTable!!!")
	}
}
