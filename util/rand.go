package util

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandom(len int) string {
	b := make([]byte, len)
	_, _ = rand.Read(b)

	return base64.URLEncoding.EncodeToString(b)
}
