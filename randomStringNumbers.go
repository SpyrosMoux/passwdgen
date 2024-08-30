package passwdgen

import (
	"math/rand"

	"github.com/SpyrosMoux/passwdgen/constants"
)

// RandomStringNumbersSymbols returns a randomized string containing numbers.
// (n) defaults to 16.
func RandomStringNumbers(n int) string {
	b := make([]byte, n)
	lettersNumbers := constants.Letters + constants.Numbers
	for i := range b {
		b[i] = lettersNumbers[rand.Intn(len(lettersNumbers))]
	}
	return string(b)
}
