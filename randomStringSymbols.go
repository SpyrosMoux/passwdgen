package passwdgen

import (
	"math/rand"

	"github.com/SpyrosMoux/passwdgen/constants"
)

// RandomStringNumbersSymbols returns a randomized string containing symbols.
// (n) defaults to 16.
func RandomStringSymbols(n int) string {
	b := make([]byte, n)
	lettersSymbols := constants.Letters + constants.Symbols
	for i := range b {
		b[i] = lettersSymbols[rand.Intn(len(lettersSymbols))]
	}
	return string(b)
}
