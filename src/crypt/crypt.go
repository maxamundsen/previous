package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"github.com/minio/highwayhash"
)

func QuickFileHash(filepath string) (string, error) {
	key := []byte("01234567890123456789012345678901")

	file, err := os.Open(filepath)
	if err != nil {
		log.Println("Error opening style.css:", err)
		return "", err
	}
	defer file.Close()

	hasher, err := highwayhash.New(key)
	if err != nil {
		log.Println("Error generating hasher.")
		return "", err
	}

	_, err = io.Copy(hasher, file)
	if err != nil {
		log.Println("Error hashing file:", err)
		return "", err
	}

	hash := hasher.Sum(nil)

	return base64.StdEncoding.EncodeToString(hash), nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func ComparePasswords(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RandBase64String(entropyBytes int) string {
	b := make([]byte, entropyBytes)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func EncryptSecret(data []byte, passKey string) (string, error) {
	key := make([]byte, 32)
	copy(key, passKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	encryptedData := gcm.Seal(nonce, nonce, data, nil)
	encodedData := base64.StdEncoding.EncodeToString(encryptedData)

	return encodedData, nil
}

func DecryptSecret(encryptedData string, passKey string) ([]byte, error) {
	encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}

	key := make([]byte, 32)
	copy(key, passKey)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedBytes) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}
	nonce, encryptedBytes := encryptedBytes[:nonceSize], encryptedBytes[nonceSize:]

	decryptedData, err := gcm.Open(nil, nonce, encryptedBytes, nil)
	if err != nil {
		return nil, err
	}

	return decryptedData, nil
}
