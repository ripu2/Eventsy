package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e *Event) Save() error {
	//add logic to store in DB
	events = append(events, *e)
	return nil
}

func GetAllEvents() ([]Event, error) {
	//add lto fetch from DB
	return events, nil
}
