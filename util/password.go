package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPasswored, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPasswored), nil
}

func CheckPassword(password, hashedPassword string) error {
	return bcrypt.
		CompareHashAndPassword([]byte(hashedPassword), []byte(password))

}
