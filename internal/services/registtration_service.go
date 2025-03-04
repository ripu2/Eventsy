package services

import (
	"errors"

	"example.com/event-management/internal/models"
)

func GetRegistrations() ([]models.Registrations, error) {
	registrations, err := models.GetRegistrationsByEvent()
	if err != nil {
		return nil, err // return early if there is an error fetching registrations. No need to continue with the rest of the function.
	}
	return registrations, nil
}

func RegisterToEvent(event *models.Event, userId int64) error {
	err := event.Register(userId)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func WithdrawRegistration(event *models.Event, userId int64) error {
	//todo
	return nil
}
