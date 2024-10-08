package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/SpyrosMoux/passwdgen"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	lengthPtr := flag.Int("length", 16, "integer indicating the length of the password")
	symbolsPtr := flag.Bool("symbols", false, "boolean indicating the use of symbols in the password")
	numbersPtr := flag.Bool("numbers", false, "boolean indicating the use of numbers in the password")

	flag.Parse()

	options := passwdgen.NewRandomStringOptions()
	options.Length = *lengthPtr

	switch *symbolsPtr {
	case true:
		switch *numbersPtr {
		case true:
			fmt.Println(passwdgen.RandomStringNumbersSymbols(&options))
		case false:
			fmt.Println(passwdgen.RandomStringSymbols(&options))
		}
	case false:
		switch *numbersPtr {
		case true:
			fmt.Println(passwdgen.RandomStringNumbers(&options))
		case false:
			fmt.Println(passwdgen.RandomString(&options))
		}
	}
}
