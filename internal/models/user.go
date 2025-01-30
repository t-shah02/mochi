package models

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type User struct {
	username      string
	password      string
	commandAccess []string
}

func NewUser(username string, password string, commandAccess string, isExisting bool) *User {
	finalPassword := ""
	if isExisting {
		finalPassword = password
	} else {
		hashedPassword, err := HashPassword(password)
		if err != nil {
			return nil
		}

		finalPassword = hashedPassword
	}

	return &User{
		username:      username,
		password:      finalPassword,
		commandAccess: strings.Split(commandAccess, ","),
	}
}
