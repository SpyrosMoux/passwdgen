package passwdgen

import (
	"math/rand"
)

const (
	Letters       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers       = "0123456789"
	Symbols       = "!@#$%^&*()_+-=~`,./:{}[];|'"
	DefaultLength = 16
)

type RandomStringOptions struct {
	Length int
}

func NewRandomStringOptions() RandomStringOptions {
	return RandomStringOptions{
		Length: 16,
	}
}

// RandomString returns a randomized string.
// Length defaults to 16.
func RandomString(options *RandomStringOptions) string {
	b := make([]byte, options.Length)
	for i := range b {
		b[i] = Letters[rand.Intn(len(Letters))]
	}
	return string(b)
}

// RandomStringNumbers returns a randomized string containing numbers.
// Length defaults to 16.
func RandomStringNumbers(options *RandomStringOptions) string {
	b := make([]byte, options.Length)
	lettersNumbers := Letters + Numbers
	for i := range b {
		b[i] = lettersNumbers[rand.Intn(len(lettersNumbers))]
	}
	return string(b)
}

// RandomStringSymbols returns a randomized string containing symbols.
// Length defaults to 16.
func RandomStringSymbols(options *RandomStringOptions) string {
	b := make([]byte, options.Length)
	lettersSymbols := Letters + Symbols
	for i := range b {
		b[i] = lettersSymbols[rand.Intn(len(lettersSymbols))]
	}
	return string(b)
}

// RandomStringNumbersSymbols returns a randomized string containing numbers and symbols.
// Length defaults to 16.
func RandomStringNumbersSymbols(options *RandomStringOptions) string {
	b := make([]byte, options.Length)
	lettersNumbersSymbols := Letters + Numbers + Symbols
	for i := range b {
		b[i] = lettersNumbersSymbols[rand.Intn(len(lettersNumbersSymbols))]
	}
	return string(b)
}
