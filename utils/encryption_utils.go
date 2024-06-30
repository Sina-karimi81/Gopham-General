package utils

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hashedPass), err
}

func CheckPasswordHash(hashedPass, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	return err == nil
}
