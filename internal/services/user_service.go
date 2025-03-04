package services

import (
	"errors"
	"time"

	"example.com/event-management/internal/models"
)

func CreateUser(user models.User) error {
	user.CreatedAt = time.Now()
	err := user.SaveUser()
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func GetUsers() ([]models.User, error) {
	users, err := models.GetAllUsers()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return users, nil
}

func LoginUser(login models.LoginRequest) (string, error) {
	token, err := login.ValidateCredentials()
	if err != nil {
		return "", errors.New(err.Error())
	}

	return token, nil
}
