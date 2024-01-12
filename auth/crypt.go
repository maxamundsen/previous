package auth

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func randBase64String(entropyBytes int) string {
	b := make([]byte, entropyBytes)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
