package utils

import (
	"crypto/rand"
	"math/big"
	"time"
)

const (
	// Characters used for generating short codes
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// GenerateShortCode generates a random short code of specified length
func GenerateShortCode(length int) string {
	result := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := range result {
		randomIndex, _ := rand.Int(rand.Reader, charsetLen)
		result[i] = charset[randomIndex.Int64()]
	}

	return string(result)
}

// GenerateID generates a unique ID based on timestamp and random string
func GenerateID() string {
	timestamp := time.Now().UnixNano()
	randomPart := GenerateShortCode(4)
	return string(rune(timestamp)) + randomPart
}

// IsValidURL performs basic URL validation
func IsValidURL(url string) bool {
	if len(url) == 0 {
		return false
	}
	
	// Basic validation - check if it starts with http:// or https://
	return len(url) > 7 && (url[:7] == "http://" || url[:8] == "https://")
}