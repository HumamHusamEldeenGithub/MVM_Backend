package utils

import "github.com/google/uuid"

// GenerateID generates a new UUID and returns it as a string
func GenerateID() string {
	return uuid.New().String()
}
