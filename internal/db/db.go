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
	createUsersTable()
	createTables()
}

func createUsersTable() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name TEXT NOT NULL,
    userName TEXT UNIQUE NOT NULL,  -- unique identifier for each user
    email TEXT UNIQUE NOT NULL,
    passwordHash TEXT NOT NULL,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL
);
`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Failed to create userTable!!! + error: " + err.Error())
	}
}

func createTables() {
	createEventTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		name TEXT NOT NULL,
    description TEXT NOT NULL,
		location TEXT NOT NULL,
    dateTime DATETIME NOT NULL,
		userId INTEGER NOT NULL,
		FOREIGN KEY(userId) REFERENCES users(id)
	)
	`
	_, err := DB.Exec(createEventTable)
	if err != nil {
		panic("Failed to create EventsTable!!! + error: " + err.Error())
	}
}
