package random

import (
	"math/rand"
	"time"
)

func NewRandomString(length uint8) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]rune, length)
	for i := range b {
		b[i] = chars[random.Intn(len(chars))]
	}

	return string(b)
}
