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

func CreateEvent(event *models.Event, creatorId int64) error {
	event.UserID = creatorId
	err := event.Save()
	if err != nil {
		return errors.New("failed to create event") // return early if there is an error saving the event. No need to continue with the rest of the function.
	}
	return nil
}

func GetEventById(eventId int64) (models.Event, error) {
	event, err := models.GetEventById(eventId)
	if err != nil {
		return models.Event{}, errors.New(err.Error())
	}
	return *event, nil
}

func UpdateEvent(event *models.Event, id int64, ownerId int64) error {
	event.ID = id
	event.UserID = ownerId
	fetchedEvent, err := GetEventById(id)
	if err != nil {
		return errors.New("event not found")
	}
	if ownerId != fetchedEvent.UserID {
		return errors.New("Unauthorized")
	}
	err = event.UpdateEvent(id)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func DeleteEvent(event *models.Event, eventId int64, ownerId int64) error {
	event.ID = eventId
	selectedEvent, err := GetEventById(eventId)
	if err != nil {
		return errors.New("event not found")
	}
	if ownerId != selectedEvent.UserID {
		return errors.New("Unauthorized")
	}
	err = selectedEvent.DeleteEvent(eventId)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
