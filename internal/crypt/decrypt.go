package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

// Decrypt decrypts an encrypted string using AES-GCM and a password.
func Decrypt(encryptedText, password string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", fmt.Errorf("invalid base64 encoding: %w", err)
	}

	if len(data) < 16+12 {
		return "", errors.New("invalid encrypted data")
	}

	salt := data[:16]
	nonce := data[16:28]
	ciphertext := data[28:]

	key, err := DeriveKey(password, salt)
	if err != nil {
		return "", fmt.Errorf("key derivation failed: %w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("AES initialization failed: %w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("GCM initialization failed: %w", err)
	}

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", errors.New("decryption failed (wrong password?)")
	}

	return string(plaintext), nil
}
