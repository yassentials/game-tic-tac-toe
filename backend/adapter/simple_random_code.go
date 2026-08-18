package adapter

import (
	"math/rand/v2"
)

type SimpleRandomCode struct {
	Length int
}

func (this SimpleRandomCode) Generate() string {
	const charsets = "QWERTYUIOPASDFGHJKLZXCVBNM"

	result := make([]byte, this.Length)

	for i := range result {
		randIndex := rand.IntN(len(charsets))
		result[i] = charsets[randIndex]
	}

	return string(result)
}
