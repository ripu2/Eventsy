package models

import (
	"errors"
	"time"

	"example.com/event-management/internal/db"
	"example.com/event-management/internal/utils"
)

type User struct {
	ID           int64
	Name         string    `binding:"required" gorm:"not null"`
	UserName     string    `binding:"required" gorm:"unique;not null"`
	Email        string    `binding:"required,email" gorm:"unique;not null"`
	PasswordHash string    `binding:"required" gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}

type LoginRequest struct {
	UserName     string `binding:"required" gorm:"unique;not null"`
	PasswordHash string `binding:"required" gorm:"not null"`
}

func (user *User) SaveUser() error {
	query := `
	INSERT INTO users(name, userName, email, passwordHash) 
	VALUES (?, ?, ?, ?)`

	processedQuery, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New(err.Error())
	}
	defer processedQuery.Close()
	hashedPassword, _ := utils.HashPassword(user.PasswordHash)
	res, err := processedQuery.Exec(user.Name, user.UserName, user.Email, hashedPassword)
	if err != nil {
		return errors.New(err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return errors.New(err.Error())
	}

	user.ID = id
	return nil
}

func GetAllUsers() ([]User, error) {
	query := `SELECT * FROM users`
	response, err := db.DB.Query(query)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer response.Close()

	var users = []User{}
	for response.Next() {
		var user User
		err := response.Scan(&user.ID, &user.Name, &user.UserName, &user.Email, &user.PasswordHash, &user.CreatedAt)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserByUserName(userName string) (User, error) {
	query := `SELECT * FROM users WHERE userName= ?`
	processedQuery, err := db.DB.Prepare(query)
	if err != nil {
		return User{}, errors.New(err.Error())
	}
	defer processedQuery.Close()

	response := processedQuery.QueryRow(userName)

	var user User
	err = response.Scan(&user.ID, &user.Name, &user.UserName, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return User{}, errors.New(err.Error())
	}
	return user, nil
}

func (u LoginRequest) ValidateCredentials() (string, error) {
	retrievedUser, err := GetUserByUserName(u.UserName)
	if err != nil {
		return "", errors.New(err.Error())
	}
	if !utils.CheckPasswordHash(u.PasswordHash, retrievedUser.PasswordHash) {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(retrievedUser.Name, u.UserName, retrievedUser.Email, retrievedUser.ID)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return token, nil
}
