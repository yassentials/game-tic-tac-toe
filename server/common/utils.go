package common

import (
	"math/rand/v2"
)

func GenRandomCode(length int) string {
	// ate the whole keyboard lol
	const charsets = "QWERTYUIOPASDFGHJKLZXCVBNM"

	result := make([]byte, length)

	for i := range result {
		randIndex := rand.IntN(len(charsets))
		result[i] = charsets[randIndex]
	}

	return string(result)
}
