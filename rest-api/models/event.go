package models

import (
	"time"

	"github.com/BaGorK/lets-go/rest-api/db"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64     `binding:"required"`
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
	insert into events (name, description, location, datetime, user_id)
	values (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.Id = id

	return err
}

func GetAllEvents() []Event {
	return events
}
