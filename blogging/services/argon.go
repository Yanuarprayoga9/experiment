package services

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/argon2"
)

func GenerateSalt(length uint32) []byte {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatalf("Failed to generate salt: %v", err)
	}
	return salt
}

func HashPassword(password string, salt []byte) string {
	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	saltBase64 := base64.RawStdEncoding.EncodeToString(salt)
	hashBase64 := base64.RawStdEncoding.EncodeToString(hash)
	return fmt.Sprintf("%s$%s", saltBase64, hashBase64)
}

func VerifyPassword(password, encodedHash string) bool {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 2 {
		return false
	}
	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false
	}
	expectedHash := base64.RawStdEncoding.EncodeToString(argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32))
	return expectedHash == parts[1]
}
