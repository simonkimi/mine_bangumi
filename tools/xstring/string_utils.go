package xstring

import (
	"math/rand"
	"strings"
)

func RemoveEmpty(iter []string) []string {
	result := make([]string, 0, len(iter))
	for _, v := range iter {
		if !IsEmptyOrWhitespace(v) {
			result = append(result, strings.TrimSpace(v))
		}
	}
	return result
}

func IsEmptyOrWhitespace(value string) bool {
	return strings.TrimSpace(value) == ""
}

func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
