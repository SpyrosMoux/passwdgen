package utils

import (
	"math/rand"
	"github.com/spyrosmoux/passwdgen/constants"
)

func RandomStringSymbols(n int) string {
	b := make([]byte, n)
	lettersSymbols := constants.Letters + constants.Symbols
	for i := range b {
		b[i] = lettersSymbols[rand.Intn(len(lettersSymbols))]
	}
	return string(b)
}