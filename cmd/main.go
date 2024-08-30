package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/SpyrosMoux/passwdgen/generators"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	lengthPtr := flag.Int("length", 16, "integer indicating the length of the password")
	symbolsPtr := flag.Bool("symbols", false, "boolean indicating the use of symbols in the password")
	numbersPtr := flag.Bool("numbers", false, "boolean indicating the use of numbers in the password")

	flag.Parse()

	switch *symbolsPtr {
	case true:
		switch *numbersPtr {
		case true:
			fmt.Println(generators.RandomStringNumbersSymbols(*lengthPtr))
		case false:
			fmt.Println(generators.RandomStringSymbols(*lengthPtr))
		}
	case false:
		switch *numbersPtr {
		case true:
			fmt.Println(generators.RandomStringNumbers(*lengthPtr))
		case false:
			fmt.Println(generators.RandomString(*lengthPtr))
		}
	}
}
