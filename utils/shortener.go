package utils

import (
	"crypto/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateShortCode(length int) string {
	b := make([]byte, length)

	_,err 	:= rand.Read(b)
	if err != nil {
		return ""
	}

	for i := 0; i < length; i++ {
		b[i] = charset[int(b[i])%len(charset)]
	}
	return  string(b)
}