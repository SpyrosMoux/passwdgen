package generators

import (
	"math/rand"

	"github.com/SpyrosMoux/passwdgen/constants"
)

func RandomStringSymbols(n int) string {
	b := make([]byte, n)
	lettersSymbols := constants.Letters + constants.Symbols
	for i := range b {
		b[i] = lettersSymbols[rand.Intn(len(lettersSymbols))]
	}
	return string(b)
}
