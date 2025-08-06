// Package utils provides utility functions for hashing passwords.
package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPass), err
}
