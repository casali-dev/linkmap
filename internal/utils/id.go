package utils

import (
	"crypto/rand"
)

const alphabet = "2346789abcdefghijkmnpqrtwxyzABCDEFGHJKLMNPQRTUVWXYZ"

func GenerateID(n int) (string, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	for i := range b {
		b[i] = alphabet[int(b[i])%len(alphabet)]
	}

	return string(b), nil
}
