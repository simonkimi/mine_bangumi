package test

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	hash, _ := bcrypt.GenerateFromPassword([]byte("awegaweg"), bcrypt.DefaultCost)
	t.Log(string(hash))
}
