package helper

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func GenerateRandomString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var builder strings.Builder

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			panic(err)
		}
		builder.WriteByte(chars[n.Int64()])
	}

	return builder.String()
}
