package generators

import (
	"math/rand"

	"github.com/SpyrosMoux/passwdgen/constants"
)

func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = constants.Letters[rand.Intn(len(constants.Letters))]
	}
	return string(b)
}
