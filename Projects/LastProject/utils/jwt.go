package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		// Check if the signing method is HMAC
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return errors.New("token is not valid")
	}

	_, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("could not parse token claims")
	}

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)

	return nil
}
