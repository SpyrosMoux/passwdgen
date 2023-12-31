package main

import (
  "fmt"
  "flag"
  "math/rand"
  "time"
  "github.com/spyrosmoux/passwdgen/utils"
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
      fmt.Println(utils.RandomStringNumbersSymbols(*lengthPtr))
    case false:
      fmt.Println(utils.RandomStringSymbols(*lengthPtr))
    }
  case false:
    switch *numbersPtr {
    case true:
      fmt.Println(utils.RandomStringNumbers(*lengthPtr))
    case false:
      fmt.Println(utils.RandomString(*lengthPtr))
    }
  }
}
