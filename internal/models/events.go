package models

import (
	"errors"

	"example.com/event-management/internal/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    string `binding:"required"`
	UserID      int64
}

func (e *Event) Save() error {
	//add logic to store in DB
	query := `
	INSERT INTO events(name, description, location, datetime, userID)
	VALUES (?, ?, ?, ?, ?)`
	processedQuery, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New(err.Error())
	}
	defer processedQuery.Close()
	res, err := processedQuery.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return errors.New(err.Error())
	}
	id, er := res.LastInsertId()
	if er != nil {
		return errors.New(er.Error())
	}
	e.ID = id
	return nil
}

func GetAllEvents() ([]Event, error) {
	//add lto fetch from DB
	query := `SELECT * FROM EVENTS`
	response, err := db.DB.Query(query)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer response.Close()
	var events = []Event{}
	for response.Next() {
		var event Event
		err := response.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `SELECT * from EVENTS WHERE ID= ? `
	processedQuery, err := db.DB.Prepare(query)
	if err != nil {
		panic(err.Error())
	}
	defer processedQuery.Close()

	response := processedQuery.QueryRow(id)

	var event Event
	err = response.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &event, nil

}

func (event *Event) UpdateEvent(id int64) error {
	query := `UPDATE EVENTS
					SET name= ? , description= ? , location= ?, datetime= ?
					WHERE id = ?`

	processedQuery, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New(err.Error())
	}
	defer processedQuery.Close()

	result, err := processedQuery.Exec(event.Name, event.Description, event.Location, event.DateTime, id)
	if err != nil {
		return errors.New(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New(err.Error())
	}
	if rowsAffected == 0 {
		return errors.New("Event not found")
	}
	return nil
}

func (event Event) DeleteEvent(id int64) error {
	query := `DELETE  FROM EVENTS WHERE id = ?`

	processedQuery, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New(err.Error())
	}
	defer processedQuery.Close()
	_, err = processedQuery.Exec(id)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}
