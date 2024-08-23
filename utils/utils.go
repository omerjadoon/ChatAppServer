package utils

import (
	"math/rand"
	"time"
)

func GenerateRoomId() string {
	// Implement a function to generate a unique room ID

	rand.Seed(time.Now().UnixNano())

	charset := "0123456789"
	result := make([]byte, 6)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}

	return string(result)
}
