package services

import (
	"errors"

	"example.com/event-management/internal/models"
)

func GetEvents() ([]models.Event, error) {
	events, err := models.GetAllEvents()
	if err != nil {
		return nil, err // return early if there is an error fetching events. No need to continue with the rest of the function.
	}
	return events, nil
}

func CreateEvent(event models.Event) error {
	event.ID = 1
	event.UserID = 1
	err := event.Save()
	if err != nil {
		return errors.New("failed to create event!!!") // return early if there is an error saving the event. No need to continue with the rest of the function.
	}
	return nil
}
