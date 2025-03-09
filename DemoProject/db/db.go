package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	// Initialize the database
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
	// Create the events table
}

func createTables() {
	// Create the users table
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create user tables")
	}
	// Create the events table
	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create event tables")
	}
	createRegistrationTable := `CREATE TABLE IF NOT EXISTS registrations (

		id INTEGER PRIMARY KEY AUTOINCREMENT,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY(event_id) REFERENCES events(id),
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`
	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		panic("Could not create registration tables")
	}

}
