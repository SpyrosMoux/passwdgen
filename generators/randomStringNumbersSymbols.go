package generators

import (
	"math/rand"

	"github.com/SpyrosMoux/passwdgen/constants"
)

func RandomStringNumbersSymbols(n int) string {
	b := make([]byte, n)
	lettersNumbersSymbols := constants.Letters + constants.Numbers + constants.Symbols
	for i := range b {
		b[i] = lettersNumbersSymbols[rand.Intn(len(lettersNumbersSymbols))]
	}
	return string(b)
}
