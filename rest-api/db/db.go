// Package db provides functions to initialize and manage the SQLite database for the API.
package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	DBL, err := sql.Open("sqlite3", "./api.db")
	if err != nil {
		panic("Could not connect to the database: " + err.Error())
	}

	DB = DBL

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTableQuery := `
	create table if not exists users (
		id integer primary key autoincrement,
		email text not null unique,
		password text not null
	)
	`
	_, err := DB.Exec(createUsersTableQuery)
	if err != nil {
		panic("could not create users table." + err.Error())
	}

	createEventsTableQuery := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		user_id INTEGER,
	  foreign key(user_id) references users(id) on delete cascade
	)
	`

	_, err = DB.Exec(createEventsTableQuery)
	if err != nil {
		panic("Could not create events table: " + err.Error())
	}
}
