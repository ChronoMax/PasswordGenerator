package generation

import (
	"math/rand"
	"time"
)

// Generates the password
func GeneratePassword(amountOfChar int) string {

	//Gets the time
	rand.Seed(time.Now().UnixNano())

	// create a byte slice with length n
	b := make([]byte, amountOfChar)

	// fill the byte slice with random ASCII values
	for i := range b {
		b[i] = byte(rand.Intn(126-33) + 33)
	}

	return string(b)
}
