package models

import (
	"errors"

	"example.com/event-management/internal/db"
)

type Registrations struct {
	ID      int64
	UserId  int64
	EventId int64
}

func GetRegistrationsByEvent() ([]Registrations, error) {
	query := "SELECT * FROM registration"
	response, err := db.DB.Query(query)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer response.Close()
	var registrations = []Registrations{}
	for response.Next() {
		var registration Registrations
		err := response.Scan(&registration.ID, &registration.UserId, &registration.EventId)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		registrations = append(registrations, registration)
	}
	return registrations, nil
}

func (event *Event) Register(userId int64) error {
	query := `INSERT INTO registration(eventId, userId) VALUES(?, ?)`

	processedQuery, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New(err.Error())
	}
	defer processedQuery.Close()
	_, err = processedQuery.Exec(event.ID, userId)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (event *Event) WithdrawRegistration(id int64) error {
	query := `DELETE FROM registration WHERE id =?`

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
