package utils

import (
  "math/rand"
  "github.com/spyrosmoux/passwdgen/constants"
)

func RandomString(n int) string {
  b := make([]byte, n)
  for i := range b {
    b[i] = constants.Letters[rand.Intn(len(constants.Letters))]
  }
  return string(b)
}
