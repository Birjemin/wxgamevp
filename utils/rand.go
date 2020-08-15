package utils

import (
	"encoding/hex"
	"math/rand"
)

// GenerateBillNo generate billNo
func GenerateBillNo() string {
	return Hex(16)
}

// Hex ...
func Hex(n int) string {
	return hex.EncodeToString(Bytes(n))
}

// Bytes generates n random bytes
func Bytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}
