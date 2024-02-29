package utils

import (
	"github.com/google/uuid"
)

// NewUUID generates a new UUID string.
func NewUUID() (string, error) {
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
