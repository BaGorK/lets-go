package models

import "time"

type Event struct {
	Id          string    `binding:"required"`
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      string    `binding:"required"`
}

var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
