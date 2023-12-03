package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateRandomHash(length int) (string, error) {
	randomBytes := make([]byte, length>>1)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(randomBytes), nil
}
