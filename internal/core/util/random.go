package util

import (
	"crypto/rand"
	"math/big"
	"monero-shop-api/internal/exception"
)

func RandomString(length uint32) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charsetLength := big.NewInt(int64(len(charset)))

	randomString := make([]byte, length)
	for i := range randomString {
		randomIndex, err := rand.Int(rand.Reader, charsetLength)
		exception.PanicLogging(err)
		randomString[i] = charset[randomIndex.Int64()]
	}

	return string(randomString)
}
