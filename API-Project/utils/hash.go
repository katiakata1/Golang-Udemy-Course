package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// byte function []byte(password) converst string to byte slice
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	// bytes will be in bytes so we need to convert it to string like we did below
	return string(bytes), err
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
