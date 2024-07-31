package db

import (
	"database/sql"
	//_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	email TEXT UNIQUE NOT NULL,
    	password TEXT NOT NULL
	)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(err)
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
   		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location    TEXT NOT NULL,
		date_time    DATETIME NOT NULL,
		user_id      INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(err)
	}

	createRegistrationsTable := `
		CREATE TABLE IF NOT EXISTS registrations (
   		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id      INTEGER,
		event_id      INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
		FOREIGN KEY(event_id) REFERENCES events(id)
	)
	`
	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic(err)
	}
}
