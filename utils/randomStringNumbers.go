package utils

import (
	"math/rand"

	"github.com/SpyrosMoux/passwdgen/constants"
)

func RandomStringNumbers(n int) string {
	b := make([]byte, n)
	lettersNumbers := constants.Letters + constants.Numbers
	for i := range b {
		b[i] = lettersNumbers[rand.Intn(len(lettersNumbers))]
	}
	return string(b)
}
