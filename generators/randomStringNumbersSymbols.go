package generators

import (
	"math/rand"

	"github.com/SpyrosMoux/passwdgen/constants"
)

// RandomStringNumbersSymbols returns a randomized string containing numbers and symbols.
// (n) defaults to 16.
func RandomStringNumbersSymbols(n int) string {
	b := make([]byte, n)
	lettersNumbersSymbols := constants.Letters + constants.Numbers + constants.Symbols
	for i := range b {
		b[i] = lettersNumbersSymbols[rand.Intn(len(lettersNumbersSymbols))]
	}
	return string(b)
}
