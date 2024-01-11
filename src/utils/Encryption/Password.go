package encryption

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

func HashPassword(pw string) (string, string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", "", nil
	}

	hash := argon2.IDKey([]byte(pw), salt, 1, 64*1024, 4, 32)
	encodedHash := base64.StdEncoding.EncodeToString(hash)
	encodedSalt := base64.StdEncoding.EncodeToString(salt)
	return encodedHash, encodedSalt, nil
}
