package crypt

import (
	"golang.org/x/crypto/scrypt"
)

// DeriveKey generates an AES key from a password using scrypt.
func DeriveKey(password string, salt []byte) ([]byte, error) {
	return scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)
}
