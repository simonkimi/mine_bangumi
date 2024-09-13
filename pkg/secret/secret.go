package secret

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomKey(len int) string {
	key := make([]byte, len)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return base64.URLEncoding.EncodeToString(key)
}
