package common

import (
	"math/rand/v2"

	"github.com/yassentials/game-tic-tac-toe/server/domain"
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

func GetRandomCharacter() domain.Character {
	availableChar := []domain.Character{domain.CHAR_O, domain.CHAR_X}

	return availableChar[rand.IntN(len(availableChar))]
}
