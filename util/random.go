package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())

}

func RandomInt(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

func RandomString(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(len(alphabet))]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "TRY"}
	return currencies[rand.Intn(len(currencies))]
}
