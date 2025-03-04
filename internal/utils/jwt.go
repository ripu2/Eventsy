package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Load secret from env
func GenerateToken(name string, userName string, email string, id int64) (string, error) {
	jwtExpires, _ := strconv.ParseInt(os.Getenv("JWT_EXPIRY_HOURS"), 10, 64)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":     name,
		"userName": userName,
		"email":    email,
		"id":       id,
		"exp":      time.Now().Add(time.Hour * time.Duration(jwtExpires)).Unix(),
	})
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", errors.New(err.Error())
	}
	return signedToken, err
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token signature")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return errors.New(err.Error())
	}

	_, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok || !parsedToken.Valid {
		return errors.New("invalid token")
	}

	// name, _ := claims["name"].(string)
	// userName, _ := claims["userName"].(string)
	// email, _ := claims["email"].(string)
	// id, _ := claims["id"].(int64)

	return nil
}
